package repo

import (
	"context"

	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/service"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"gorm.io/gorm"
)

type roleRepo struct {
	q *query.Query
}

var _ service.IRoleRepo = (*roleRepo)(nil)

func NewRoleRepo(db *gorm.DB) service.IRoleRepo {
	return &roleRepo{q: query.Use(db)}
}

func (repo *roleRepo) Find(ctx context.Context, req *v1.RoleListRequest) (result []*model.Role, count int64, err error) {
	q := repo.q.Role
	return q.WithContext(ctx).FindByPage(req.Offset(), req.Limit())
}

func (repo *roleRepo) FindOne(ctx context.Context, req *v1.RoleRequest) (*model.Role, error) {
	q := repo.q.Role
	return q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
}

func (repo *roleRepo) Create(ctx context.Context, data *model.Role) error {
	return repo.q.Role.WithContext(ctx).Create(data)
}

func (repo *roleRepo) Update(ctx context.Context, req *v1.RoleUpdateRequest) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *roleRepo) UpdateStatus(ctx context.Context, id int64, status bool) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Update(q.Status, status)
	return err
}

func (repo *roleRepo) Delete(ctx context.Context, req *v1.RoleDeleteRequest) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.In(req.IDs...)).Delete()
	return err
}
