// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameAdmin = "admin"

// Admin 后台用户表
type Admin struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string                `gorm:"column:username;not null" json:"username"`
	Password   string                `gorm:"column:password;not null" json:"-"`
	Icon       *string               `gorm:"column:icon;comment:头像" json:"icon"`                                        // 头像
	Email      *string               `gorm:"column:email;comment:邮箱" json:"email"`                                      // 邮箱
	NickName   *string               `gorm:"column:nick_name;comment:昵称" json:"nickName"`                               // 昵称
	Note       *string               `gorm:"column:note;comment:备注信息" json:"note"`                                      // 备注信息
	LoginTime  *time.Time            `gorm:"column:login_time;comment:最后登录时间" json:"loginTime"`                         // 最后登录时间
	Status     bool                  `gorm:"column:status;not null;default:1;comment:帐号启用状态：0->禁用；1->启用" json:"status"` // 帐号启用状态：0->禁用；1->启用
	IsDel      soft_delete.DeletedAt `gorm:"column:is_del;not null" json:"-"`
	CreateTime *time.Time            `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"`
	UpdateTime *time.Time            `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"updateTime"`
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}