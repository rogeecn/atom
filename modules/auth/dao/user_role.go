package dao

import (
	"atom/database/models"
	"atom/database/query"
	"context"
)

type UserRoleDao interface {
	Exists(context.Context, int) bool
	Create(context.Context, int, int) error
	Update(context.Context, int, int) error
	Delete(context.Context, int, int) error
}

type userRoleDaoImpl struct {
	query *query.Query
}

func NewUserRoleDao(query *query.Query) UserRoleDao {
	return &userRoleDaoImpl{query: query}
}

func (dao *userRoleDaoImpl) Exists(ctx context.Context, userID int) bool {
	userRole := dao.query.UserRole
	count, _ := userRole.WithContext(ctx).Where(userRole.UserID.Eq(uint64(userID))).Count()
	return count > 0
}

func (dao *userRoleDaoImpl) Create(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	return userRole.WithContext(ctx).Create(&models.UserRole{
		UserID: uint64(userID),
		RoleID: uint64(roleID),
	})
}

func (dao *userRoleDaoImpl) Update(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	_, err := userRole.WithContext(ctx).Where(userRole.UserID.Eq(uint64(userID))).Update(userRole.RoleID, roleID)
	return err
}

func (dao *userRoleDaoImpl) Delete(ctx context.Context, userID, roleID int) error {
	userRole := dao.query.UserRole
	_, err := userRole.WithContext(ctx).
		Where(userRole.UserID.Eq(uint64(userID))).
		Where(userRole.RoleID.Eq(uint64(roleID))).
		Delete()
	return err
}
