package dao

import (
	"atom/database/models"
	"atom/database/query"
	"context"
)

type UserRoleDao struct {
	query *query.Query
}

func NewUserRoleDao(query *query.Query) *UserRoleDao {
	return &UserRoleDao{query: query}
}

func (dao *UserRoleDao) Exists(ctx context.Context, userID int) bool {
	userRole := dao.query.UserRole
	count, _ := userRole.WithContext(ctx).Where(userRole.UserID.Eq(uint64(userID))).Count()
	return count > 0
}

func (dao *UserRoleDao) Create(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	return userRole.WithContext(ctx).Create(&models.UserRole{
		UserID: uint64(userID),
		RoleID: uint64(roleID),
	})
}

func (dao *UserRoleDao) Update(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	_, err := userRole.WithContext(ctx).Where(userRole.UserID.Eq(uint64(userID))).Update(userRole.RoleID, roleID)
	return err
}

func (dao *UserRoleDao) Delete(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	_, err := userRole.WithContext(ctx).
		Where(userRole.UserID.Eq(uint64(userID))).
		Where(userRole.RoleID.Eq(uint64(roleID))).
		Delete()
	return err
}
