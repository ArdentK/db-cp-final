package postgres

import (
	"context"
	"database/sql"

	"github.com/ArdentK/db-cp-final/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.QueryRow(
		"INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id",
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID)
}

func (r *UserRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	u := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = $1 AND password = $2",
		email,
		password,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}
