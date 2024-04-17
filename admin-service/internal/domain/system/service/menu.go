package service

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/domain/system/dto"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/app/validator"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/page"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type IMenuRepo interface {
	Find(ctx context.Context, req *v1.MenuListRequest) ([]*model.Menu, int64, error)
	FindOne(ctx context.Context, req *v1.MenuRequest) (*model.Menu, error)
	Create(ctx context.Context, data *model.Menu) error
	Update(ctx context.Context, req *v1.MenuUpdateRequest) error
	UpdateHidden(ctx context.Context, ID int64, hidden *bool) error
	Delete(ctx context.Context, req *v1.MenuDeleteRequest) error
}

type menuService struct {
	repo IMenuRepo
}

var _ v1.IMenuServer = (*menuService)(nil)

func NewMenuService(repo IMenuRepo) v1.IMenuServer {
	return &menuService{repo: repo}
}

func (svc *menuService) List(ctx context.Context, req *v1.MenuListRequest) (*page.Data[*model.Menu], error) {
	list, total, err := svc.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return &page.Data[*model.Menu]{
		Total: total,
		List:  list,
	}, nil
}

func (svc *menuService) TreeList(ctx context.Context) ([]*dto.MenuNode, error) {
	req := v1.MenuListRequest{Query: page.NewNot()}
	list, _, err := svc.repo.Find(ctx, &req)
	if err != nil {
		return nil, err
	}
	return svc.listToTree(list), nil
}

func (svc *menuService) GetDetails(ctx context.Context, req *v1.MenuRequest) (*model.Menu, error) {
	if err := validator.StructCtx(ctx, req); err != nil {
		return nil, errs.BadRequest.WithDetails(err)
	}
	return svc.repo.FindOne(ctx, req)
}

func (svc *menuService) Add(ctx context.Context, req *v1.MenuCreateRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	var data model.Menu
	copier.Copy(&data, req)
	return svc.repo.Create(ctx, &data)
}

func (svc *menuService) Update(ctx context.Context, req *v1.MenuUpdateRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	return svc.repo.Update(ctx, req)
}

func (svc *menuService) UpdateHidden(ctx context.Context, req *v1.MenuUpdateHiddenRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	return svc.repo.UpdateHidden(ctx, req.ID, req.Hidden)
}

func (svc *menuService) Delete(ctx context.Context, req *v1.MenuDeleteRequest) error {
	if err := validator.StructCtx(ctx, req); err != nil {
		return errs.BadRequest.WithDetails(err)
	}
	return svc.repo.Delete(ctx, req)
}

func (svc *menuService) listToTree(list []*model.Menu) []*dto.MenuNode {
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
