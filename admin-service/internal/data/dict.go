package data

import (
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IDictRepo interface {
	// CreateDict 创建字典
	CreateDict(c *model.SysDict) (*model.SysDict, error)
	// GetDict 获取一个字典通过categoryCode和dictCode
	GetDict(categoryCode, dictCode uint) (*model.SysDict, error)
	// UpdateDict 更新字典
	UpdateDict(c *model.SysDict) error
	// DeleteDict 删除字典通过categoryCode和dictCode
	DeleteDict(categoryCode, dictCode uint) error
	// ListDictByCategoryCode 查找某一组字典集合
	ListDictByCategoryCode(categoryCode uint) ([]*model.SysDict, error)
	// ListDict 查找满足条件的字典
	ListDict(c *model.SysDict) ([]*model.SysDict, error)
}

type dDictRepo struct {
	db *gorm.DB
}

func NewDictRepo(db *gorm.DB) IDictRepo {
	return &dDictRepo{db}
}

func (d dDictRepo) CreateDict(c *model.SysDict) (*model.SysDict, error) {
	err := d.db.Create(c).Error
	return c, errors.WithStack(err)
}

func (d dDictRepo) GetDict(categoryCode, dictCode uint) (*model.SysDict, error) {
	sysDict := new(model.SysDict)
	err := d.db.Where("category_code = ? AND dictCode = ?", categoryCode, dictCode).First(sysDict).Error
	return sysDict, errors.WithStack(err)
}

func (d dDictRepo) UpdateDict(c *model.SysDict) error {
	if c.CategoryCode == nil || c.Code == nil {
		return errors.Errorf("bad param, categoryCode=%d, code=%d", *c.CategoryCode, *c.Code)
	}
	// update by categoryCode and code
	m := &model.SysDict{CategoryCode: c.CategoryCode, Code: c.Code}
	err := d.db.Model(m).Updates(c).Error
	return errors.WithStack(err)
}

func (d dDictRepo) DeleteDict(categoryCode, dictCode uint) error {
	m := &model.SysDict{CategoryCode: &categoryCode, Code: &dictCode}
	err := d.db.Delete(m).Error
	return errors.WithStack(err)
}

func (d dDictRepo) ListDictByCategoryCode(categoryCode uint) ([]*model.SysDict, error) {
	var dicts []*model.SysDict
	err := d.db.Where("categoryCode == ?", categoryCode).Find(dicts).Error
	return dicts, errors.WithStack(err)
}

func (d dDictRepo) ListDict(c *model.SysDict) ([]*model.SysDict, error) {
	var dicts []*model.SysDict
	err := d.db.Where(c).Find(dicts).Error
	return dicts, errors.WithStack(err)
}