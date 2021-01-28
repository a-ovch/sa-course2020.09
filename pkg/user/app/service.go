package app

import "github.com/a-ovch/sa-course2020.09/pkg/user/domain"

type Service struct {
	ur domain.UserRepository
}

func (s *Service) CreateUser() *domain.User {
	return nil
}

func NewService(ur domain.UserRepository) *Service {
	return &Service{ur: ur}
}
