package repo

import (
	"context"
	"github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/query"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
	q  *query.Query
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db: db, q: query.Use(db)}
}

func (repo *RoleRepo) Find(ctx context.Context, req *v1.RoleListRequest) (result []*model.Role, count int64, err error) {
	q := repo.q.Role
	return q.WithContext(ctx).FindByPage(req.Offset(), req.Limit())
}

func (repo *RoleRepo) FindOne(ctx context.Context, req *v1.RoleRequest) (*model.Role, error) {
	q := repo.q.Role
	return q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
}

func (repo *RoleRepo) Create(ctx context.Context, data *model.Role) error {
	return repo.q.Role.WithContext(ctx).Create(data)
}

func (repo *RoleRepo) Update(ctx context.Context, req *v1.RoleUpdateRequest) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *RoleRepo) UpdateStatus(ctx context.Context, id int64, status bool) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Update(q.Status, status)
	return err
}

func (repo *RoleRepo) Delete(ctx context.Context, req *v1.RoleDeleteRequest) error {
	q := repo.q.Role
	_, err := q.WithContext(ctx).Where(q.ID.In(req.IDs...)).Delete()
	return err
}