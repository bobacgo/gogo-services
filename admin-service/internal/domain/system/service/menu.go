package service

import (
	"context"

	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/dto"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type IMenuRepo interface {
	Find(ctx context.Context, req *v1.MenuListRequest) ([]*model.Menu, int64, error)
	FindOne(ctx context.Context, req *v1.MenuRequest) (*model.Menu, error)
	Create(ctx context.Context, data *model.Menu) error
	Update(ctx context.Context, req *v1.MenuUpdateRequest) error
	UpdateHidden(ctx context.Context, ID int64, hidden bool) error
	Delete(ctx context.Context, req *v1.MenuDeleteRequest) error
}

type MenuService struct {
	repo IMenuRepo
}

func NewMenuService(repo IMenuRepo) *MenuService {
	return &MenuService{repo: repo}
}

func (svc *MenuService) List(ctx context.Context, req *v1.MenuListRequest) (*page.Data[*model.Menu], error) {
	list, total, err := svc.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return &page.Data[*model.Menu]{
		Total: total,
		List:  list,
	}, nil
}

func (svc *MenuService) TreeList(ctx context.Context) ([]*dto.MenuNode, error) {
	req := v1.MenuListRequest{
		Query: page.Query{
			PageNum:  -1,
			PageSize: -1,
		},
	}
	list, _, err := svc.repo.Find(ctx, &req)
	if err != nil {
		return nil, err
	}
	return svc.listToTree(list), nil
}

func (svc *MenuService) GetDetails(ctx context.Context, req *v1.MenuRequest) (*v1.MenuResponse, error) {
	one, err := svc.repo.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.MenuResponse{
		Menu: one,
	}, nil
}

func (svc *MenuService) Add(ctx context.Context, req *v1.MenuCreateRequest) error {
	var data model.Menu
	copier.Copy(&data, req)
	return svc.repo.Create(ctx, &data)
}

func (svc *MenuService) Update(ctx context.Context, req *v1.MenuUpdateRequest) error {
	return svc.repo.Update(ctx, req)
}

func (svc *MenuService) UpdateHidden(ctx context.Context, ID int64, hidden bool) error {
	return svc.repo.UpdateHidden(ctx, ID, hidden)
}

func (svc *MenuService) Delete(ctx context.Context, req *v1.MenuDeleteRequest) error {
	return svc.repo.Delete(ctx, req)
}

func (svc *MenuService) listToTree(list []*model.Menu) []*dto.MenuNode {
	tree := make([]*dto.MenuNode, 0, len(list))
	pMap := lo.GroupBy(list, func(e *model.Menu) int64 {
		return e.ParentID
	})
	for _, v := range list {
		mn := &dto.MenuNode{Menu: *v}
		mn.Children = pMap[v.ID]
		tree = append(tree, mn)
	}
	return lo.Filter(tree, func(e *dto.MenuNode, idx int) bool {
		return e.ParentID == 0
	})
}
