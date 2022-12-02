package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// SysUser 用户表
type SysUser struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"unique,comment:用户账号"`
	Password  string `gorm:"not null"`
	Nickname  string `gorm:"comment:用户昵称"`
	Gender    uint   `gorm:"default:2,comment:性别 0-女|1-男|2-保密"`
	Email     string `gorm:"comment:邮箱"`
	PhoneNo   string `gorm:"unique,comment:手机号"`
	Avatar    string `gorm:"comment:头像标识"`
	Status    uint8  `gorm:"default:0,comment:是否启用 0-正常|1-停用"`
	Remark    string `gorm:"comment:备注"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDel     soft_delete.DeletedAt `orm:"softDelete:flag,comment:0-未删除|1-已删除"`
}