package biz

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/internal/data"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
)

type IDict interface {
	// CreateDict 创建字典
	CreateDict(ctx context.Context, c *model.SysDict) (*r.Resp, error)
	// UpdateDict 更新字典
	UpdateDict(ctx context.Context, c *model.SysDict) (*r.Resp, error)
	// DeleteDict 删除字典通过categoryCode和dictCode
	DeleteDict(ctx context.Context, categoryCode, dictCode uint32) (*r.Resp, error)

	// GetDictMap 查找满足条件的字典
	GetDictMap(ctx context.Context, c *model.SysDict) (*r.Resp, error)
}

type bDict struct{}

var DictBiz IDict = bDict{}

func (b bDict) CreateDict(ctx context.Context, c *model.SysDict) (*r.Resp, error) {
	repo := data.NewDictRepo(g.DB.WithContext(ctx))
	dict := &model.SysDict{}
	_, _ = repo.CreateDict(dict)
	panic("implement me")
}

func (b bDict) UpdateDict(ctx context.Context, c *model.SysDict) (*r.Resp, error) {
	//TODO implement me
	panic("implement me")
}

func (b bDict) DeleteDict(ctx context.Context, categoryCode, dictCode uint32) (*r.Resp, error) {
	//TODO implement me
	panic("implement me")
}

func (b bDict) GetDictMap(ctx context.Context, c *model.SysDict) (*r.Resp, error) {
	//TODO implement me
	panic("implement me")
}
