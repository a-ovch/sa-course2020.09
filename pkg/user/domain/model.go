package domain

import "github.com/a-ovch/sa-course2020.09/pkg/common/domain"

type UserID domain.UUID

type User struct {
	ID        UserID
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UserRepository interface {
	NextID() UserID
	Store(user *User) error
	Find(id UserID) (*User, error)
	Delete(id UserID) error
}
