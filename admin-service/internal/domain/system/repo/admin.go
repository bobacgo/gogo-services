package repo

import (
	"context"
	"errors"
	"time"

	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/dto"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type AdminRepo struct {
	q *query.Query
}

var _ service.IAdminRepo = (*AdminRepo)(nil)

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{
		q: query.Use(db),
	}
}

func (repo *AdminRepo) FindAdminRole(ctx context.Context, adminID int64) ([]*model.Role, error) {
	qARR := repo.q.AdminRoleRelation
	q := repo.q.Role

	var roleIDs []int64
	if err := qARR.WithContext(ctx).Where(qARR.AdminID.Eq(adminID)).Pluck(qARR.RoleID, &roleIDs); err != nil {
		return nil, err
	}
	return q.WithContext(ctx).Where(q.ID.In(roleIDs...)).Find()
}

func (repo *AdminRepo) UpdateRole(ctx context.Context, adminID int64, role []int64) error {
	q := repo.q.AdminRoleRelation
	res, err := q.WithContext(ctx).Where(q.AdminID.Eq(adminID)).Delete()
	if res.RowsAffected == 0 {
		if err != nil {
			return err
		}
		return gorm.ErrRecordNotFound
	}
	var data []*model.AdminRoleRelation
	for _, v := range role {
		data = append(data, &model.AdminRoleRelation{
			AdminID: adminID,
			RoleID:  v,
		})
	}
	return q.WithContext(ctx).Create(data...)
}

func (repo *AdminRepo) HasUsername(ctx context.Context, req *dto.UniqueUsernameQuery) (*dto.UniqueResult, error) {
	q := repo.q.Admin

	conds := []gen.Condition{
		q.Username.Eq(req.Username),
	}
	if req.ExcludeID != 0 {
		conds = append(conds, q.ID.Neq(req.ExcludeID))
	}
	return repo.hasRecord(ctx, conds)
}

func (repo *AdminRepo) HasEmail(ctx context.Context, req *dto.UniqueEmailQuery) (*dto.UniqueResult, error) {
	q := repo.q.Admin

	conds := []gen.Condition{
		q.Email.Eq(req.Email),
	}
	if req.ExcludeID != 0 {
		conds = append(conds, q.ID.Neq(req.ExcludeID))
	}
	return repo.hasRecord(ctx, conds)
}

func (repo *AdminRepo) FindByUsername(ctx context.Context, username string) (*model.Admin, error) {
	q := repo.q.Admin
	return q.WithContext(ctx).Where(q.Username.Eq(username)).First()
}

func (repo *AdminRepo) FindByID(ctx context.Context, ID int64) (*model.Admin, error) {
	q := repo.q.Admin
	return q.WithContext(ctx).Where(q.ID.Eq(ID)).First()
}

func (repo *AdminRepo) Find(ctx context.Context, req *v1.AdminListRequest) (*page.Data[*model.Admin], error) {
	q := repo.q.Admin
	do := q.WithContext(ctx)
	if req.Keyword != "" {
		do = do.Or(q.Username.Like("%" + req.Keyword + "%")).
			Or(q.Nickname.Like("%" + req.Keyword + "%"))
	}
	list, count, err := do.FindByPage(req.Offset(), req.Limit())
	if err != nil {
		return nil, err
	}
	return page.New(count, list...), nil
}

func (repo *AdminRepo) Insert(ctx context.Context, records ...*model.Admin) error {
	return repo.q.Admin.WithContext(ctx).Create(records...)
}

func (repo *AdminRepo) Update(ctx context.Context, data *v1.AdminUpdateRequest) error {
	q := repo.q.Admin
	_, err := repo.q.Admin.WithContext(ctx).Where(q.ID.Eq(data.ID)).Updates(data)
	return err
}

func (repo *AdminRepo) UpdateLoginTime(ctx context.Context, ID int64, loginTime time.Time) error {
	q := repo.q.Admin
	_, err := q.WithContext(ctx).Where(q.ID.Eq(ID)).Update(q.LoginTime, loginTime)
	return err
}

func (repo *AdminRepo) UpdatePwd(ctx context.Context, ID int64, pwd string) error {
	q := repo.q.Admin
	_, err := q.WithContext(ctx).Where(q.ID.Eq(ID)).Update(q.Password, pwd)
	return err
}

func (repo *AdminRepo) UpdateStatus(ctx context.Context, ID int64, status bool) error {
	q := repo.q.Admin
	_, err := q.WithContext(ctx).Where(q.ID.Eq(ID)).Update(q.Status, status)
	return err
}

func (repo *AdminRepo) Delete(ctx context.Context, ID int64) error {
	q := repo.q.Admin
	_, err := q.WithContext(ctx).Where(q.ID.Eq(ID)).Delete()
	return err
}

func (repo *AdminRepo) hasRecord(ctx context.Context, conds []gen.Condition) (*dto.UniqueResult, error) {
	if len(conds) == 0 {
		return nil, errors.New("conds is empty")
	}
	admin, err := repo.q.Admin.WithContext(ctx).Unscoped().Select(repo.q.Admin.IsDel).Where(conds...).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &dto.UniqueResult{ // 数据不存在
				Exist: false,
				IsDel: 0,
			}, nil
		}
		return nil, err // 查找出错
	}
	return &dto.UniqueResult{
		Exist: true,
		IsDel: uint8(admin.IsDel),
	}, nil
}
