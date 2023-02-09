package service

import (
	"atom/database/models"
	"atom/modules/auth/dao"
	"atom/modules/auth/dto"
	"atom/providers/jwt"
	"context"
)

type UserService struct {
	userRoleDao dao.UserRoleDao
	userDao     dao.UserDao
	jwt         *jwt.JWT
}

func NewUserService(
	userRoleDao dao.UserRoleDao,
	userDao dao.UserDao,
	jwt *jwt.JWT,
) *UserService {
	return &UserService{
		userRoleDao: userRoleDao,
		userDao:     userDao,
		jwt:         jwt,
	}
}

func (svc *UserService) AttachRole(ctx context.Context, userID, roleID int) error {
	if svc.userRoleDao.Exists(ctx, userID) {
		return svc.userRoleDao.Update(ctx, userID, roleID)
	}
	return svc.userRoleDao.Create(ctx, userID, roleID)
}

func (svc *UserService) DetachRole(ctx context.Context, userID, roleID int) error {
	if !svc.userRoleDao.Exists(ctx, userID) {
		return nil
	}
	return svc.userRoleDao.Delete(ctx, userID, roleID)
}
func (svc *UserService) AuthMatchPassword(ctx context.Context, req *dto.LoginRequestForm) (*models.User, error) {
	return &models.User{
		ID:       10,
		UUID:     "1",
		Username: "2",
		Password: "3",
		Nickname: "4",
		Avatar:   "5",
		RoleID:   66,
	}, nil
}

func (svc *UserService) GenerateJWTTokenFromUser(ctx context.Context, user *models.User) (string, error) {
	return svc.jwt.CreateToken(svc.jwt.CreateClaims(jwt.BaseClaims{
		UID:  user.ID,
		Role: user.RoleID,
	}))
}
