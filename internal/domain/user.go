package domain

import "time"

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Role         Role
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
