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
	Label           string         `gorm:"column:label;type:varchar(191)" json:"label"`                       // 展示值
	Value           int64          `gorm:"column:value;type:bigint(20)" json:"value"`                         // 字典值
	Status          bool           `gorm:"column:status;type:tinyint(1)" json:"status"`                       // 启用状态
	Sort            int64          `gorm:"column:sort;type:bigint(20)" json:"sort"`                           // 排序标记
	SysDictionaryID int64          `gorm:"column:sys_dictionary_id;type:bigint(20)" json:"sys_dictionary_id"` // 关联标记
}

// TableName SysDictionaryDetail's table name
func (*SysDictionaryDetail) TableName() string {
	return TableNameSysDictionaryDetail
}
