package usecase

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

type jwt interface {
	GenerateToken(userID uuid.UUID) (string, error)
}

type repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}
