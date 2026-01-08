package service

import (
	"errors"
	"movie-reservation-system/internal/domain"
	"movie-reservation-system/internal/repository"
	"movie-reservation-system/internal/security"

	"github.com/google/uuid"
)
type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService{
	return &AuthService{userRepo: userRepo}
}


func (s *AuthService) Signup(name, email, password string) error {
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}
	hash, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID: uuid.NewString(),
		Name: name,
		Email: email,
		PasswordHash: hash,
		Role: domain.Role{
			ID: 2,
		},
		IsActive: true,
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string)(*domain.User, error){
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !security.CheckPassword(password, user.PasswordHash){
		return nil, errors.New("Invalid credentials")
	}
	return user, nil
}