package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleEmployee  = "employee"
	RoleModerator = "moderator"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
