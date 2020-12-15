package service

import (
	"fmt"
	"project/internal/http/dao"
)

func New(dao dao.Dao) *Service {
	return &Service{dao: dao}
}

type Service struct {
	dao dao.Dao
}

func (s *Service) ProcessBusiness() {
	res := s.dao.GetUserInfo()
	fmt.Println(res)
}
