// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRole = "role"

// Role 后台用户角色表
type Role struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string     `gorm:"column:name;not null;uniqueIndex:role_unique,priority:1;comment:名称" json:"name"` // 名称
	Description *string    `gorm:"column:description;comment:描述" json:"description"`                               // 描述
	AdminCount  int32      `gorm:"column:admin_count;not null;comment:后台用户数量" json:"adminCount"`                   // 后台用户数量
	CreateTime  *time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`                              // 创建时间
	Status      bool       `gorm:"column:status;not null;default:1;comment:启用状态：0->禁用；1->启用" json:"status"`        // 启用状态：0->禁用；1->启用
	Sort        *int32     `gorm:"column:sort" json:"sort"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}