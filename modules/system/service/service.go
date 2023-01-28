package service

import (
	"atom/modules/system/dao"
	"atom/modules/system/dto"
	"context"
)

type SystemService interface {
	GetName(ctx context.Context) (dto.Name, error)
}

type systemService struct {
	dao dao.Dao
}

func NewSystemService(dao dao.Dao) SystemService {
	return &systemService{dao: dao}
}

func (svc *systemService) GetName(ctx context.Context) (dto.Name, error) {
	if err := svc.dao.Release(ctx, 10, "Rogee"); err != nil {
		return dto.Name{}, err
	}
	return dto.Name{Name: "System.GetName"}, nil
}
