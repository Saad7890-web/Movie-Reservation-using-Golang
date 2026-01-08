package postgres

import (
	"database/sql"
	"errors"
	"movie-reservation-system/internal/domain"
	"movie-reservation-system/internal/repository"
)

type userRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password_hash, role_id)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Name
		user.Email
		user.PasswordHash
		user.Role.ID

	)

	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.password_hash, u.is_active,
		       r.id, r.name
		FROM users u
		JOIN roles r ON r.id = u.role_id
		WHERE u.email = $1
	`

	row := r.db.QueryRow(query, email)

	user := &domain.User{}
	user.Role = domain.Role{}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email
		&user.PasswordHash,
		&user.IsActive,
		&user.Role.ID,
		&user.Role.Name,
	)

	if err == sql.ErrNoRows{
		return nil, errors.New("usr not found")
	}
	return user, err
}

