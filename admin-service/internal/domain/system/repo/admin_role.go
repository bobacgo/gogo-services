package repo

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"gorm.io/gorm"
)

type adminRoleRepo struct {
	q *query.Query
}

var _ service.IAdminRoleRepo = (*adminRoleRepo)(nil)

func NewAdminRoleRepo(db *gorm.DB) service.IAdminRoleRepo {
	return &adminRoleRepo{q: query.Use(db)}
}

func (repo *adminRoleRepo) FindAdminRole(ctx context.Context, adminID int64) ([]*model.Role, error) {
	qARR := repo.q.AdminRoleRelation
	q := repo.q.Role

	var roleIDs []int64
	if err := qARR.WithContext(ctx).Where(qARR.AdminID.Eq(adminID)).Pluck(qARR.RoleID, &roleIDs); err != nil {
		return nil, err
	}
	return q.WithContext(ctx).Where(q.ID.In(roleIDs...)).Find()
}

func (repo *adminRoleRepo) UpdateRole(ctx context.Context, adminID int64, role []int64) error {
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
