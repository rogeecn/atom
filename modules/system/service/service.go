package service

import (
	"atom/modules/system/dao"
)

type SystemService interface {
}

type systemService struct {
	dao dao.Dao
}

func NewSystemService(dao dao.Dao) SystemService {
	return &systemService{dao: dao}
}
