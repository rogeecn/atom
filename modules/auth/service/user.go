package service

import (
	"atom/modules/auth/dao"
	"context"
)

type UserService interface {
	AttachRole(context.Context, int, int) error
}

type userService struct {
	userRoleDao dao.UserRoleDao
	userDao     dao.UserDao
}

func NewUserService(
	userRoleDao dao.UserRoleDao,
	userDao dao.UserDao,
) UserService {
	return &userService{
		userRoleDao: userRoleDao,
		userDao:     userDao,
	}
}

func (svc *userService) AttachRole(ctx context.Context, userID, roleID int) error {
	if svc.userRoleDao.Exists(ctx, userID) {
		return svc.userRoleDao.Update(ctx, userID, roleID)
	}
	return svc.userRoleDao.Create(ctx, userID, roleID)
}

func (svc *userService) DetachRole(ctx context.Context, userID, roleID int) error {
	if !svc.userRoleDao.Exists(ctx, userID) {
		return nil
	}
	return svc.userRoleDao.Delete(ctx, userID, roleID)
}
