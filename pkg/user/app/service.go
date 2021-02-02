package app

import (
	"errors"

	"github.com/a-ovch/sa-course2020.09/pkg/user/domain"
)

var ErrUserNotFound = errors.New("user not found")

type UserData struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserDataWithID struct {
	ID string `json:"id"`
	UserData
}

type Service struct {
	ur domain.UserRepository
}

func (s *Service) CreateUser(d *UserData) (*domain.User, error) {
	u := &domain.User{
		ID:        s.ur.NextID(),
		Username:  d.Username,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		Phone:     d.Phone,
	}

	err := s.ur.Store(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) FindUser(id string) (*UserDataWithID, error) {
	uid := domain.UserID(id)
	u, err := s.ur.Find(uid)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrUserNotFound
	}

	return &UserDataWithID{
		string(u.ID),
		UserData{
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		},
	}, nil
}

func (s *Service) UpdateUser(id string, data *UserData) error {
	u, err := s.ur.Find(domain.UserID(id))
	if err != nil {
		return err
	}

	if u == nil {
		return ErrUserNotFound
	}

	u.Username = data.Username
	u.FirstName = data.FirstName
	u.LastName = data.LastName
	u.Email = data.Email
	u.Phone = data.Phone

	return s.ur.Store(u)
}

func (s *Service) DeleteUser(id string) error {
	return s.ur.Delete(domain.UserID(id))
}

func NewService(ur domain.UserRepository) *Service {
	return &Service{ur: ur}
}
