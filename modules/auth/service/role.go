package service

import (
	"context"
)

type RoleService interface {
	Create(ctx context.Context) error
}

type roleService struct {
}

func NewRoleService() RoleService {
	return &roleService{}
}

func (svc *roleService) Create(ctx context.Context) error {
	return nil
}
