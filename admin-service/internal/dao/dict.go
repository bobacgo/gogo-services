package dao

import (
	"context"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/common-lib/g"
	"github.com/pkg/errors"
)

type IDictRepo interface {
	// CreateDict 创建字典
	CreateDict(ctx context.Context, sysDict *model.SysDict) (ID uint, err error)
	// GetDict 获取一个字典通过categoryCode和dictCode
	GetDict(ctx context.Context, categoryCode, dictCode uint) (*model.SysDict, error)
	// UpdateDict 更新字典 by categoryCode and dictCode
	UpdateDict(ctx context.Context, sysDict *model.SysDict) error
	// DeleteDict 删除字典通过categoryCode和dictCode
	DeleteDict(ctx context.Context, categoryCode, dictCode *uint) error
	// ListDictByCategoryCode 查找某一组字典集合
	ListDictByCategoryCode(ctx context.Context, categoryCode uint) ([]*model.SysDict, error)
	// ListDict 查找满足条件的字典
	ListDict(ctx context.Context, sysDict *model.SysDict) ([]*model.SysDict, error)
}

type dictRepo struct{}

func (dictRepo) CreateDict(ctx context.Context, m *model.SysDict) (uint, error) {
	err := g.DB.WithContext(ctx).Create(m).Error
	return m.ID, errors.WithStack(err)
}

func (dictRepo) GetDict(ctx context.Context, categoryCode, dictCode uint) (*model.SysDict, error) {
	sysDict := new(model.SysDict)
	err := g.DB.WithContext(ctx).Where("category_code = ? AND dictCode = ?", categoryCode, dictCode).First(sysDict).Error
	return sysDict, errors.WithStack(err)
}

func (dictRepo) UpdateDict(ctx context.Context, m *model.SysDict) error {
	if m.CategoryCode == nil || m.Code == nil {
		return errors.Errorf("bad param, categoryCode=%d, code=%d", *m.CategoryCode, *m.Code)
	}
	// update by categoryCode and code
	sd := &model.SysDict{CategoryCode: m.CategoryCode, Code: m.Code}
	err := g.DB.WithContext(ctx).Model(sd).Updates(m).Error
	return errors.WithStack(err)
}

func (dictRepo) DeleteDict(ctx context.Context, categoryCode, dictCode *uint) error {
	m := &model.SysDict{CategoryCode: categoryCode, Code: dictCode}
	err := g.DB.WithContext(ctx).Delete(m).Error
	return errors.WithStack(err)
}

func (dictRepo) ListDictByCategoryCode(ctx context.Context, categoryCode uint) ([]*model.SysDict, error) {
	var list []*model.SysDict
	err := g.DB.WithContext(ctx).Where("categoryCode == ?", categoryCode).Find(list).Error
	return list, errors.WithStack(err)
}

func (dictRepo) ListDict(ctx context.Context, sysDict *model.SysDict) ([]*model.SysDict, error) {
	var list []*model.SysDict
	err := g.DB.WithContext(ctx).Where(sysDict).Find(list).Error
	return list, errors.WithStack(err)
}
