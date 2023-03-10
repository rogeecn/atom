// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDictionaryDetail = "sys_dictionary_details"

// SysDictionaryDetail mapped from table <sys_dictionary_details>
type SysDictionaryDetail struct {
	ID              uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	SysDictionaryID int64          `gorm:"column:sys_dictionary_id;type:bigint(20)" json:"sys_dictionary_id"` // 关联标记
	Label           string         `gorm:"column:label;type:varchar(191)" json:"label"`                       // 展示值
	Value           string         `gorm:"column:value;type:varchar(191)" json:"value"`                       // 字典值
	Status          bool           `gorm:"column:status;type:tinyint(1)" json:"status"`                       // 启用状态
	Weight          int64          `gorm:"column:weight;type:bigint(20)" json:"weight"`                       // 排序权重
}

// TableName SysDictionaryDetail's table name
func (*SysDictionaryDetail) TableName() string {
	return TableNameSysDictionaryDetail
}
