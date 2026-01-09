package repository

import "movie-reservation-system/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	UpdateRole(userID string, roleName string) error
}

