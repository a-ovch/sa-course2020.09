package infrastructure

import (
	"database/sql"
	"github.com/google/uuid"

	"github.com/a-ovch/sa-course2020.09/pkg/common/infrastructure/database"
	"github.com/a-ovch/sa-course2020.09/pkg/user/domain"
)

type userRepository struct {
	client database.Client
}

func (r *userRepository) NextID() domain.UserID {
	return domain.UserID(uuid.New().String())
}

func (r *userRepository) Store(u *domain.User) error {
	const query = "INSERT INTO \"user\" (id, username, first_name, last_name, email, phone) " +
		"VALUES ($1, $2, $3, $4, $5, $6) " +
		"ON CONFLICT (id) DO UPDATE " +
		"SET username = $2, first_name = $3, last_name = $4, email = $5, phone = $6"
	return r.client.Exec(query, u.ID, u.Username, u.FirstName, u.LastName, u.Email, u.Phone)
}

func (r *userRepository) Find(id domain.UserID) (*domain.User, error) {
	const query = `SELECT id, username, first_name, last_name, email, phone FROM "user" WHERE id = $1`

	u := new(domain.User)

	row := r.client.QueryRow(query, id)
	err := row.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) Delete(id domain.UserID) error {
	const query = `DELETE FROM "user" WHERE id = $1`
	return r.client.Exec(query, id)
}

func NewUserRepository(client database.Client) domain.UserRepository {
	return &userRepository{client: client}
}
