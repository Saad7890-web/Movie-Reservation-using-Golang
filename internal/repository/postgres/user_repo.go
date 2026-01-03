package postgres

import (
	"context"
	"database/sql"

	"github.com/saad7890/movie-reservation/internal/domain"
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return  &UserRepository{db: db}
}

func(r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	return r.db.QueryRowContext(
		ctx,
		`INSERT INTO users (email, password, role)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`,
		u.Email, u.Password, u.Role,
	).Scan(&u.ID, &u.CreatedAt)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := &domain.User{}
	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, email, password, role, created_at
		 FROM users WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	u := &domain.User{}
	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, email, password, role, created_at
		 FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)

	if err != nil {
		return nil, err
	}
	return u, nil
}