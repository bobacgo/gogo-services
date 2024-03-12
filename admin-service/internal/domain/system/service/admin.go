package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/config"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/app/security"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uid"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/utime"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IAdminRoleRepo interface {
	FindAdminRole(ctx context.Context, adminID int64) ([]*model.Role, error)
	UpdateRole(ctx context.Context, adminID int64, role []int64) error
}

type IAdminRepo interface {
	// HasUsername
	// 1.查询字段少。
	// 2.不能通过
	HasUsername(ctx context.Context, username string) (exist bool, isDel uint8, err error)
	HasEmail(ctx context.Context, email string) (exist bool, isDel uint8, err error)
	FindByUsername(ctx context.Context, username string) (*model.Admin, error)
	FindByID(ctx context.Context, ID int64) (*model.Admin, error)
	Find(ctx context.Context, req *v1.ListRequest) (*page.Data[*model.Admin], error)
	Insert(ctx context.Context, record ...*model.Admin) error
	Update(ctx context.Context, data *model.Admin) error
	UpdatePwd(ctx context.Context, ID int64, pwd string) error
	UpdateStatus(ctx context.Context, ID int64, status bool) error
	Delete(ctx context.Context, ID int64) error
}

const (
	LoginLimitKeyPrefix = "admin:login_limit"
)

type AdminService struct {
	cache         redis.Cmdable
	repo          IAdminRepo
	adminRoleRepo IAdminRoleRepo
}

func NewAdminService(rdb redis.Cmdable, repo IAdminRepo, adminRoleRepo IAdminRoleRepo) *AdminService {
	return &AdminService{
		cache:         rdb,
		repo:          repo,
		adminRoleRepo: adminRoleRepo,
	}
}

func (svc *AdminService) Register(ctx context.Context, data *v1.AdminRegisterRequest) error {
	// 查询是否有相同的用户名
	exist, isDel, err := svc.repo.HasUsername(ctx, data.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 其他错误(非用户未找到)
		return err
	}
	if isDel == 1 { // 已注销
		return errs.AdminUnUsernameDuplicated
	}
	if exist {
		return errs.AdminUsernameDuplicated
	}
	exist, isDel, err = svc.repo.HasEmail(ctx, data.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 其他错误(非邮箱未找到)
		return err
	}
	if exist && isDel == 0 { // 邮箱已存在(不包括已注销的)
		// 重新启用已注销的账户, 需要校验邮箱是否重复,如果重复就清空该账户的邮箱号.
		return errs.AdminEmailDuplicated
	}
	return svc.repo.Insert(ctx, &model.Admin{
		Username: data.Username,
		Password: data.Password.BcryptHash(),
		Icon:     &data.Icon,
		Email:    &data.Email,
		Note:     &data.Note,
	})
}

func (svc *AdminService) Login(ctx context.Context, data *v1.AdminLoginRequest) (*v1.AdminLoginResponse, error) {
	// 1.支持多平台， 不同平台有不同的 token
	// 2.每一个平台只能登录同时在线一个
	admin, err := svc.repo.FindByUsername(ctx, data.Username)
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
	if !pwdHelper.BcryptVerifyWithCount(ctx, admin.Password, string(data.Password)) {
		return nil, errs.AdminLoginFail.WithDetails(v1.AdminPwdErr{DecrCount: pwdHelper.GetRemainCount()})
	}
	claims := svc.newClaims(admin)
	atoken, rtoken, err := security.JwtHelper.Generate(ctx, claims)
	if err != nil {
		return nil, errs.TokenGenerateErr
	}
	return &v1.AdminLoginResponse{Token: atoken, RToken: rtoken}, nil
}

func (svc *AdminService) Logout(ctx context.Context, username string) error {
	return security.JwtHelper.RemoveToken(ctx, username)
}

func (svc *AdminService) RefreshToken(ctx context.Context, req *v1.AdminRefreshTokenRequest) (*v1.AdminLoginResponse, error) {
	claims, err := security.JwtHelper.Parse(req.AToken)
	if !security.JwtHelper.ValidationErrorExpired(err) {
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

func (svc *AdminService) GetAdminInfo(ctx context.Context, username string) (*v1.UserInfo, error) {
	return &v1.UserInfo{
		Username: username,
		NickName: "nickName",
		Roles:    []string{"admin"},
		Menus:    "menus",
		Icon:     "icon",
	}, nil
}

func (svc *AdminService) List(ctx context.Context, req *v1.ListRequest) (*page.Data[*model.Admin], error) {
	return svc.repo.Find(ctx, req)
}

func (svc *AdminService) GetItem(ctx context.Context, ID int64) (*model.Admin, error) {
	admin, err := svc.repo.FindByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.AdminNotFound
	}
	return admin, err
}

func (svc *AdminService) Update(ctx context.Context, data *model.Admin) error {
	return svc.repo.Update(ctx, data)
}

func (svc *AdminService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) error {
	admin, err := svc.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.AdminNotFound
		}
		return err
	}
	if admin.Password != string(req.Password) { // TODO
		return errs.AdminOldPwdErr
	}
	if err = svc.repo.UpdatePwd(ctx, admin.ID, string(req.NewPassword)); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *AdminService) Delete(ctx context.Context, ID int64) error {
	admin, err := svc.repo.FindByID(ctx, ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err = svc.repo.Delete(ctx, ID); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *AdminService) UpdateStatus(ctx context.Context, ID int64, status bool) error {
	admin, err := svc.repo.FindByID(ctx, ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err = svc.repo.UpdateStatus(ctx, ID, status); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *AdminService) UpdateRole(ctx context.Context, ID int64, role []int64) error {
	admin, err := svc.repo.FindByID(ctx, ID)
	if err != nil {
		return errs.AdminNotFound
	}
	if err = svc.adminRoleRepo.UpdateRole(ctx, ID, role); err != nil {
		return err
	}
	return security.JwtHelper.RemoveToken(ctx, admin.Username)
}

func (svc *AdminService) GetRoleList(ctx context.Context, ID int64) ([]*model.Role, error) {
	return svc.adminRoleRepo.FindAdminRole(ctx, ID)
}

func (svc *AdminService) newPasswdVerifier(ctx context.Context, admin *model.Admin) *security.PasswdVerifier {
	pwdHelper := security.NewPasswdVerifier(svc.cache, config.Conf.Service.ErrAttemptLimit)
	// 第二天0点清零
	remain := utime.ZeroHour(1).Unix() - time.Now().Unix()
	pwdHelper.SetKey(fmt.Sprintf("%s:%s", LoginLimitKeyPrefix, admin.Username), time.Duration(remain)*time.Second)
	pwdHelper.OnErr = func(err error) {
		if errors.Is(err, security.ErrPasswdLimit) && admin.Status {
			if err = svc.UpdateStatus(ctx, admin.ID, false); err != nil {
				logger.Error("update status:", "id", admin.ID, "err", err)
			}
			return
		}
		logger.Error("verify password err:", "username", admin.Username, "err", err)
	}
	return pwdHelper
}

func (svc *AdminService) newClaims(admin *model.Admin) *security.Claims {
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
