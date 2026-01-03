package domain

import "time"


type Role string

const(
	RoleUser Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID int64
	Email string
	Password string
	Role Role
	CreatedAt time.Time
}