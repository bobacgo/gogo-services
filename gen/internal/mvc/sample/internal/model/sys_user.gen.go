// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser 用户表
type SysUser struct {
	ID        uint32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string     `gorm:"column:username;not null;uniqueIndex:sys_user_unique,priority:1;comment:账户名" json:"username"` // 账户名
	Password  string     `gorm:"column:password;not null" json:"-"`
	PhoneNo   *string    `gorm:"column:phone_no;uniqueIndex:sys_user_unique_2,priority:1;comment:手机号" json:"phone_no"` // 手机号
	Email     *string    `gorm:"column:email;uniqueIndex:sys_user_unique_1,priority:1" json:"email"`
	Birthday  *time.Time `gorm:"column:birthday;comment:出生日期" json:"birthday"` // 出生日期
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
