package service

import (
	"context"

	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"github.com/jinzhu/copier"
)

type IRoleRepo interface {
	Find(ctx context.Context, req *v1.RoleListRequest) ([]*model.Role, int64, error)
	FindOne(ctx context.Context, req *v1.RoleRequest) (*model.Role, error)
	Create(ctx context.Context, data *model.Role) error
	Update(ctx context.Context, req *v1.RoleUpdateRequest) error
	UpdateStatus(ctx context.Context, id int64, status bool) error
	Delete(ctx context.Context, req *v1.RoleDeleteRequest) error
}

type roleService struct {
	repo IRoleRepo
}

var _ v1.IRoleServer = (*roleService)(nil)

func NewRoleService(repo IRoleRepo) v1.IRoleServer {
	return &roleService{repo: repo}
}

func (svc *roleService) List(ctx context.Context, req *v1.RoleListRequest) (*page.Data[*model.Role], error) {
	list, total, err := svc.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return &page.Data[*model.Role]{
		Total: total,
		List:  list,
	}, nil
}

func (svc *roleService) GetDetails(ctx context.Context, req *v1.RoleRequest) (*v1.RoleResponse, error) {
	one, err := svc.repo.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.RoleResponse{
		Role: one,
	}, nil
}

func (svc *roleService) Add(ctx context.Context, req *v1.RoleCreateRequest) error {
	var data model.Role
	copier.Copy(&data, req)
	return svc.repo.Create(ctx, &data)
}

func (svc *roleService) Update(ctx context.Context, req *v1.RoleUpdateRequest) error {
	return svc.repo.Update(ctx, req)
}

func (svc *roleService) Delete(ctx context.Context, req *v1.RoleDeleteRequest) error {
	return svc.repo.Delete(ctx, req)
}

func (svc *roleService) UpdateStatus(ctx context.Context, id int64, status bool) error {
	return svc.repo.UpdateStatus(ctx, id, status)
}
