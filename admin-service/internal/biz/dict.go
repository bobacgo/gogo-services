package biz

import (
	"context"
	"github.com/gogoclouds/gogo-services/admin-service/internal/dao"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
)

type IDictService interface {
	// CreateDict 创建字典
	CreateDict(ctx context.Context, c *model.SysDict) *r.Resp
	// UpdateDict 更新字典
	UpdateDict(ctx context.Context, c *model.SysDict) *r.Resp
	// DeleteDict 删除字典通过categoryCode和dictCode
	DeleteDict(ctx context.Context, categoryCode, dictCode *uint) *r.Resp
	// GetDictMap 查找满足条件的字典
	GetDictMap(ctx context.Context, categoryCodes []uint) *r.Resp
}

type dictService struct{}

func (dictService) CreateDict(ctx context.Context, m *model.SysDict) *r.Resp {
	if _, err := dao.DictRepo.CreateDict(ctx, m); err != nil {
		g.Log.Errorf("%+v", err)
		return r.FailMsg(r.CreateFail)
	}
	return r.SuccessMsg(r.CreateSuccess)
}

func (dictService) UpdateDict(ctx context.Context, m *model.SysDict) *r.Resp {
	if err := dao.DictRepo.UpdateDict(ctx, m); err != nil {
		g.Log.Errorf("%+v", err)
		return r.FailMsg(r.UpdateFail)
	}
	return r.SuccessMsg(r.UpdateSuccess)
}

func (dictService) DeleteDict(ctx context.Context, categoryCode, dictCode *uint) *r.Resp {
	if err := dao.DictRepo.DeleteDict(ctx, categoryCode, dictCode); err != nil {
		g.Log.Errorf("%+v", err)
		return r.FailMsg(r.DeleteFail)
	}
	return r.SuccessMsg(r.DeleteSuccess)
}

func (dictService) GetDictMap(ctx context.Context, categoryCodes []uint) *r.Resp {
	//TODO implement me
	panic("implement me")
}
