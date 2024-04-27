package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/dto"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/framework/app/security"
	"github.com/gogoclouds/gogo-services/framework/app/validator"
	"github.com/gogoclouds/gogo-services/framework/pkg/uid"
	"github.com/gogoclouds/gogo-services/framework/pkg/utime"
	"github.com/gogoclouds/gogo-services/framework/web/r/page"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type IAdminRoleRepo interface {
	FindAdminRole(ctx context.Context, adminID int64) ([]*model.Role, error)
	UpdateRole(ctx context.Context, adminID int64, role []int64) error
}

type IAdminRepo interface {
	// HasUsername
	// 1.查询字段少。
	// 2.不能通过
	HasUsername(ctx context.Context, req *dto.UniqueUsernameQuery) (*dto.UniqueResult, error)
	HasEmail(ctx context.Context, req *dto.UniqueEmailQuery) (*dto.UniqueResult, error)
	FindByUsername(ctx context.Context, username string) (*model.Admin, error)
	FindByID(ctx context.Context, ID int64) (*model.Admin, error)
	Find(ctx context.Context, req *v1.AdminListRequest) (*page.Data[*model.Admin], error)
	// 创建admin
	Insert(ctx context.Context, record ...*model.Admin) error
	// 更新admin相关
	Update(ctx context.Context, data *v1.AdminUpdateRequest) error
	UpdateLoginTime(ctx context.Context, ID int64, loginTime time.Time) error
	UpdatePwd(ctx context.Context, ID int64, pwd string) error
	UpdateStatus(ctx context.Context, ID int64, status bool) error
	// 删除admin
	Delete(ctx context.Context, ID int64) error
}

const (
	LoginLimitKeyPrefix = "admin:login_limit"
)

type adminService struct {
	cache         redis.Cmdable
	repo          IAdminRepo
	adminRoleRepo IAdminRoleRepo
	menuRepo      IMenuRepo
}

var _ v1.IAdminServer = (*adminService)(nil)

func NewAdminService(rdb redis.Cmdable, repo IAdminRepo, adminRoleRepo IAdminRoleRepo, menuRepo IMenuRepo) v1.IAdminServer {
	return &adminService{
		cache:         rdb,
		repo:          repo,
		adminRoleRepo: adminRoleRepo,
		menuRepo:      menuRepo,
	}
}

func (svc *adminService) Register(ctx context.Context, req *v1.AdminRegisterRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}

	// 查询是否有相同的用户名
	hasRes, err := svc.repo.HasUsername(ctx, &dto.UniqueUsernameQuery{
		Username: req.Username,
	})
	if err != nil { // 其他错误(非用户未找到)
		return err
	}
	if hasRes.IsDel == 1 { // 已注销
		return errs.AdminUnUsernameDuplicated
	}
	if hasRes.Exist {
		return errs.AdminUsernameDuplicated
	}
	hasRes, err = svc.repo.HasEmail(ctx, &dto.UniqueEmailQuery{
		Email: req.Email,
	})
	if err != nil { // 其他错误(非邮箱未找到)
		return err
	}
	if hasRes.Exist && hasRes.IsDel == 0 { // 邮箱已存在(不包括已注销的)
		// 重新启用已注销的账户, 需要校验邮箱是否重复,如果重复就清空该账户的邮箱号.
		return errs.AdminEmailDuplicated
	}
	return svc.repo.Insert(ctx, &model.Admin{
		Username: req.Username,
		Password: req.Password.BcryptHash(),
		Icon:     &req.Icon,
		Email:    &req.Email,
		Note:     &req.Note,
	})
}

func (svc *adminService) Login(ctx context.Context, req *v1.AdminLoginRequest) (*v1.AdminLoginResponse, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	// 1.支持多平台， 不同平台有不同的 token
	// 2.每一个平台只能登录同时在线一个
	admin, err := svc.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.AdminLoginFail
		}
		return nil, err
	}

	if !admin.Status {
		return nil, errs.AdminLoginForbidden
	}

	pwdHelper := svc.newPasswdVerifier(ctx, admin)
	if !pwdHelper.BcryptVerifyWithCount(ctx, admin.Password, string(req.Password)) {
		return nil, errs.AdminLoginFail.WithDetails(v1.AdminPwdErr{DecrCount: pwdHelper.GetRemainCount()})
	}
	claims := svc.newClaims(admin)
	atoken, rtoken, err := security.JwtHelper.Generate(ctx, claims)
	if err != nil {
		return nil, errs.TokenGenerateErr
	}

	go svc.repo.UpdateLoginTime(ctx, admin.ID, time.Now())

	return &v1.AdminLoginResponse{TokenHead: "Bearer ", Token: atoken, RToken: rtoken}, nil
}

func (svc *adminService) Logout(ctx context.Context, req *v1.AdminLogoutRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	return security.JwtHelper.RemoveToken(ctx, req.Username)
}

func (svc *adminService) RefreshToken(ctx context.Context, req *v1.AdminRefreshTokenRequest) (*v1.AdminLoginResponse, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	claims, err := security.JwtHelper.Parse(req.AToken)
	if err != nil && !security.JwtHelper.ValidationErrorExpired(err) {
		return nil, err
	}

	admin, err := svc.repo.FindByUsername(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	if !admin.Status {
		return nil, errs.AdminLoginForbidden
	}
	claims = svc.newClaims(admin)
	aToken, rToken, err := security.JwtHelper.Refresh(ctx, req.RToken, claims)
	if err != nil {
		return nil, err
	}
	return &v1.AdminLoginResponse{Token: aToken, RToken: rToken}, nil
}

func (svc *adminService) GetAdminInfo(ctx context.Context, req *v1.AdminInfoRequest) (*v1.UserInfo, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		slog.ErrorContext(ctx, "get admin info error", "username", req.Username, "err", err)
		return nil, errs.AdminNotFound
	}

	menusList, _, err := svc.menuRepo.Find(ctx, &v1.MenuListRequest{Query: page.NewNot()})
	if err != nil {
		slog.ErrorContext(ctx, "get menu list error", "err", err)
		return nil, errs.AdminNotFound
	}

	// menus := new(MenuService).listToTree(menusList)

	if admin.Nickname == nil {
		admin.Nickname = &admin.Username
	}
	var menus []*v1.AdminMenu
	copier.Copy(&menus, menusList)
	userInfo := &v1.UserInfo{
		Username: admin.Username,
		NickName: *admin.Nickname,
		Roles:    []string{"超级管理员"}, // TODO 获取用户角色
		Menus:    menus,             // TODO 获取角色菜单
	}

	if admin.Icon != nil {
		userInfo.Icon = *admin.Icon
	}
	return userInfo, nil
}

func (svc *adminService) List(ctx context.Context, req *v1.AdminListRequest) (*page.Data[*model.Admin], error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	return svc.repo.Find(ctx, req)
}

func (svc *adminService) GetItem(ctx context.Context, req *v1.AdminRequest) (*v1.AdminResponse, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByID(ctx, req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.AdminNotFound
	}

	return &v1.AdminResponse{
		ID:         admin.ID,
		Username:   admin.Username,
		Email:      admin.Email,
		Icon:       admin.Icon,
		Note:       admin.Note,
		Status:     admin.Status,
		LoginTime:  admin.LoginTime,
		CreateTime: admin.CreateTime,
		UpdateTime: admin.UpdateTime,
	}, err
}

// 更新admin信息
// 不是很重要的信息不需要实时更新(没有移除token)
func (svc *adminService) Update(ctx context.Context, req *v1.AdminUpdateRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	if req.Email != nil {
		hesRes, err := svc.repo.HasEmail(ctx, &dto.UniqueEmailQuery{
			ExcludeID: req.ID,
			Email:     *req.Email,
		})
		if err != nil {
			return err
		}
		if hesRes.Exist && hesRes.IsDel == 0 { // 邮箱已存在(不包括已注销的)
			return errs.AdminEmailDuplicated
		}
	}
	return svc.repo.Update(ctx, req)
}

func (svc *adminService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.AdminNotFound
		}
		return err
	}
	if !req.Password.BcryptVerify(admin.Password) {
		return errs.AdminOldPwdErr
	}
	if err = svc.repo.UpdatePwd(ctx, admin.ID, req.NewPassword.BcryptHash()); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *adminService) Delete(ctx context.Context, req *v1.AdminRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByID(ctx, req.ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err := svc.repo.Delete(ctx, req.ID); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *adminService) UpdateStatus(ctx context.Context, req *v1.AdminUpdateStatusRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByID(ctx, req.ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err = svc.repo.UpdateStatus(ctx, req.ID, *req.Status); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *adminService) UpdateRole(ctx context.Context, req *v1.AdminUpdateRoleRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	admin, err := svc.repo.FindByID(ctx, req.ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err = svc.adminRoleRepo.UpdateRole(ctx, req.ID, req.Roles); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *adminService) GetRoleList(ctx context.Context, req *v1.AdminRequest) ([]*model.Role, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	return svc.adminRoleRepo.FindAdminRole(ctx, req.ID)
}

func (svc *adminService) newPasswdVerifier(ctx context.Context, admin *model.Admin) *security.PasswdVerifier {
	pwdHelper := security.NewPasswdVerifier(svc.cache, config.Cfg.Service.ErrAttemptLimit)
	// 第二天0点清零
	remain := utime.ZeroHour(1).Unix() - time.Now().Unix()
	pwdHelper.SetKey(fmt.Sprintf("%s:%s", LoginLimitKeyPrefix, admin.Username), time.Duration(remain)*time.Second)
	pwdHelper.OnErr = func(err error) {
		if errors.Is(err, security.ErrPasswdLimit) && admin.Status {
			if err = svc.UpdateStatus(ctx, &v1.AdminUpdateStatusRequest{ID: admin.ID, Status: lo.ToPtr(false)}); err != nil {
				slog.ErrorContext(ctx, "update status:", "id", admin.ID, "err", err)
			}
			return
		}
		slog.ErrorContext(ctx, "verify password err:", "username", admin.Username, "err", err)
	}
	return pwdHelper
}

func (svc *adminService) newClaims(admin *model.Admin) *security.Claims {
	claims := security.Claims{
		StandardClaims: jwt.StandardClaims{Id: uid.UUID()},
		UserID:         strconv.FormatInt(admin.ID, 10),
		Username:       admin.Username,
		Roles:          nil, // TODO
	}
	if admin.Nickname != nil {
		claims.Nickname = *admin.Nickname
	} else {
		claims.Nickname = admin.Username
	}
	return &claims
}
