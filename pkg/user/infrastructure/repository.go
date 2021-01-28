package infrastructure

import (
	"github.com/google/uuid"

	"github.com/a-ovch/sa-course2020.09/pkg/common/infrastructure/database"
	"github.com/a-ovch/sa-course2020.09/pkg/user/domain"
)

type userRepository struct {
	client database.Client
}

func (r *userRepository) NextID() domain.UserID {
	return domain.UserID(uuid.New())
}

func (r *userRepository) Store(user *domain.User) error {
	return nil
}

func NewUserRepository(client database.Client) domain.UserRepository {
	return &userRepository{client: client}
}
