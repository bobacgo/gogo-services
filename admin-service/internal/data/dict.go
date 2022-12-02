package data

import (
	"context"

	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"gorm.io/gorm"
)

type IDictRepo interface {
	// CreateDict 创建字典
	CreateDict(ctx context.Context, c *model.SysDict) (*model.SysDict, error)
	// GetDict 获取一个字典通过categoryCode和dictCode
	GetDict(ctx context.Context, categoryCode, dictCode uint32) (*model.SysDict, error)
	// UpdateDict 更新字典
	UpdateDict(ctx context.Context, c *model.SysDict) (*model.SysDict, error)
	// DeleteDict 删除字典通过categoryCode和dictCode
	DeleteDict(ctx context.Context, categoryCode, dictCode uint32) (bool, error)
	// ListDictByCategoryCode 查找某一组字典集合
	ListDictByCategoryCode(ctx context.Context, categoryCode uint32) ([]*model.SysDict, error)
	// ListDict 查找满足条件的字典
	ListDict(ctx context.Context, c *model.SysDict) ([]*model.SysDict, error)
}

type dDictRepo struct {
	db *gorm.DB
}

func NewDictRepo(db *gorm.DB) IDictRepo {
	return &dDictRepo{db}
}

func (d dDictRepo) CreateDict(ctx context.Context, c *model.SysDict) (*model.SysDict, error) {
	//TODO implement me
	panic("implement me")
}

func (d dDictRepo) GetDict(ctx context.Context, categoryCode, dictCode uint32) (*model.SysDict, error) {
	//TODO implement me
	panic("implement me")
}

func (d dDictRepo) UpdateDict(ctx context.Context, c *model.SysDict) (*model.SysDict, error) {
	//TODO implement me
	panic("implement me")
}

func (d dDictRepo) DeleteDict(ctx context.Context, categoryCode, dictCode uint32) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (d dDictRepo) ListDictByCategoryCode(ctx context.Context, categoryCode uint32) ([]*model.SysDict, error) {
	//TODO implement me
	panic("implement me")
}

func (d dDictRepo) ListDict(ctx context.Context, c *model.SysDict) ([]*model.SysDict, error) {
	//TODO implement me
	panic("implement me")
}
