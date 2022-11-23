package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Dict 字典表
// categoryCode 和 code 组成复合主键
type Dict struct {
	CategoryCode uint32 `json:"categoryCode" orm:"primaryKey,comment:字典类型"` // 字典类型
	Code         uint32 `json:"code" orm:"primaryKey,comment:字典code"`
	Value        string `json:"value" orm:"primaryKey,comment:字典value"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsDel        soft_delete.DeletedAt `orm:"softDelete:flag,comment:0-未删除|1-已删除"`
}
