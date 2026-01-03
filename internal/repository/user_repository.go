package repository

import (
	"context"

	"github.com/saad7890/movie-reservation/internal/domain"
)


type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id int64) (*domain.User, error)
}