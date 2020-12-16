package service

import "task/http/internal/dao"

func NewUser(dao dao.UserDao) *Service {
	return &Service{userdao: dao}
}

type Service struct {
	userdao dao.UserDao
}

func (s *Service) GetUserList() string {
	res := s.userdao.GetUsers()
	return res
}
