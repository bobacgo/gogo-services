package repo

import (
	"context"
	"errors"

	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"gorm.io/gorm"
)

type AdminRepo struct {
	q *query.Query
}

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
	_, err := q.WithContext(ctx).Where(q.AdminID.Eq(adminID)).Delete()
	if err != nil {
		return err
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

func (repo *AdminRepo) HasUsername(ctx context.Context, username string) (exist bool, isDel uint8, err error) {
	q := repo.q.Admin
	admin, err := q.WithContext(ctx).Unscoped().Select(q.IsDel).Where(q.Username.Eq(username)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, 0, nil
		}
		return false, 0, err
	}
	return true, uint8(admin.IsDel), nil
}

func (repo *AdminRepo) HasEmail(ctx context.Context, email string) (exist bool, isDel uint8, err error) {
	q := repo.q.Admin
	admin, err := q.WithContext(ctx).Unscoped().Select(q.IsDel).Where(q.Email.Eq(email)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, 0, nil
		}
		return false, 0, err
	}
	return true, uint8(admin.IsDel), nil
}

func (repo *AdminRepo) FindByUsername(ctx context.Context, username string) (*model.Admin, error) {
	q := repo.q.Admin
	return q.WithContext(ctx).Where(q.Username.Eq(username)).First()
}

func (repo *AdminRepo) FindByID(ctx context.Context, ID int64) (*model.Admin, error) {
	q := repo.q.Admin
	return q.WithContext(ctx).Where(q.ID.Eq(ID)).First()
}

func (repo *AdminRepo) Find(ctx context.Context, req *v1.ListRequest) (*page.Data[*model.Admin], error) {
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

func (repo *AdminRepo) Update(ctx context.Context, data *model.Admin) error {
	_, err := repo.q.Admin.WithContext(ctx).Updates(data)
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
