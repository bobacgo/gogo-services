package repo

import (
	"context"
	v1 "github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/api/user/v1"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/model"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/query"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
	q  *query.Query
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db, q: query.Use(db)}
}

func (repo *UserRepo) Find(ctx context.Context, req *v1.UserListRequest) (result []*model.SysUser, count int64, err error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).FindByPage(req.Offset(), req.Limit())
}

func (repo *UserRepo) FindOne(ctx context.Context, req *v1.UserRequest) (*model.SysUser, error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
}

func (repo *UserRepo) Create(ctx context.Context, data *model.SysUser) error {
	return repo.q.SysUser.WithContext(ctx).Create(data)
}

func (repo *UserRepo) Update(ctx context.Context, req *v1.UserUpdateRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *UserRepo) Delete(ctx context.Context, req *v1.UserDeleteRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Delete()
	return err
}
