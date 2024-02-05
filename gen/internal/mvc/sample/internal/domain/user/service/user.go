package service

import (
	"context"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	v1 "github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/api/user/v1"
	"github.com/gogoclouds/gogo-services/gen/internal/mvc/sample/internal/model"
	"github.com/jinzhu/copier"
)

type IUserRepo interface {
	Find(ctx context.Context, req *v1.UserListRequest) ([]*model.SysUser, int64, error)
	FindOne(ctx context.Context, req *v1.UserRequest) (*model.SysUser, error)
	Create(ctx context.Context, data *model.SysUser) error
	Update(ctx context.Context, req *v1.UserUpdateRequest) error
	Delete(ctx context.Context, req *v1.UserDeleteRequest) error
}

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) List(ctx context.Context, req *v1.UserListRequest) (*page.Data[*model.SysUser], error) {
	list, total, err := svc.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return page.New(total, list...), nil
}

func (svc *UserService) GetDetails(ctx context.Context, req *v1.UserRequest) (*v1.UserResponse, error) {
	one, err := svc.repo.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.UserResponse{
		SysUser: one,
	}, nil
}

func (svc *UserService) Add(ctx context.Context, req *v1.UserCreateRequest) error {
	var data model.SysUser
	copier.Copy(&data, req)
	return svc.repo.Create(ctx, &data)
}

func (svc *UserService) Update(ctx context.Context, req *v1.UserUpdateRequest) error {
	return svc.repo.Update(ctx, req)
}

func (svc *UserService) Delete(ctx context.Context, req *v1.UserDeleteRequest) error {
	return svc.repo.Delete(ctx, req)
}
