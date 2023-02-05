// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	UUID      string         `gorm:"column:uuid;type:varchar(191)" json:"uuid"`                        // UUID
	Username  string         `gorm:"column:username;type:varchar(191)" json:"username"`                // 登录名
	Password  string         `gorm:"column:password;type:varchar(191)" json:"password"`                // 登录密码
	Nickname  string         `gorm:"column:nickname;type:varchar(191)" json:"nickname"`                // 昵称
	Avatar    string         `gorm:"column:avatar;type:varchar(191)" json:"avatar"`                    // 头像
	RoleID    uint64         `gorm:"column:role_id;type:bigint(20) unsigned;default:1" json:"role_id"` // 角色ID
	Phone     string         `gorm:"column:phone;type:varchar(191)" json:"phone"`                      // 手机号
	Email     string         `gorm:"column:email;type:varchar(191)" json:"email"`                      // 邮箱
	Status    string         `gorm:"column:status;type:varchar(191);default:ok" json:"status"`         // 用户状态
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}