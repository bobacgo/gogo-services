package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// SysDict 字典表
// categoryCode 和 code 组成复合主键
type SysDict struct {
	ID           uint   `gorm:"primaryKey,autoIncrement"` // 可以用于查询优化
	Namespace    string `json:"namespace" orm:"default:default,comment:业务领域"`
	CategoryCode *uint  `json:"categoryCode" orm:"primaryKey,comment:字典类型"`
	Code         *uint  `json:"code" orm:"primaryKey,comment:字典code"`
	Value        string `json:"value" orm:"not null,comment:字典value"`
	Remark       string `json:"remark" orm:"comment:说明"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsDel        soft_delete.DeletedAt `orm:"softDelete:flag,comment:0-未删除|1-已删除"`
}