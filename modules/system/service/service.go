package service

import (
	"atom/modules/system/dao"
)

type SystemService struct {
	dao dao.Dao
}

func NewSystemService(dao dao.Dao) *SystemService {
	return &SystemService{dao: dao}
}
