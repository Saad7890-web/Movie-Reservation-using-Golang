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
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role.ID,

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
		&user.Email,
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

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	query := `
		SELECT 
			u.id,
			u.name,
			u.email,
			u.password_hash,
			u.is_active,
			u.created_at,
			u.updated_at,
			r.id,
			r.name
		FROM users u
			u.id,
			u.name,
			u.email,
			u.password_hash,
			u.is_active,
			u.created_at,
			u.updated_at,
			r.id,
			r.name
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1
	`

	var user domain.User
	var role domain.Role

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&role.ID,
		&role.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.Role = role
	return &user, nil
}


func (r *userRepository) UpdateRole(userID string, roleName string) error {
	query := `
		UPDATE users
		SET role_id = (
			SELECT id FROM roles WHERE name = $1
		),
		updated_at = NOW()
		WHERE id = $2
	`

	result, err := r.db.Exec(query, roleName, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user or role not found")
	}

	return nil
}


