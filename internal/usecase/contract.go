package usecase

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

type jwt interface {
	GenerateToken(claims domain.Claims) (string, error)
}

type repository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	CreatePvz(ctx context.Context, pvz domain.Pvz) error
	GetLastReception(ctx context.Context, pvzID uuid.UUID) (domain.Reception, error)
	CreateReception(ctx context.Context, reception domain.Reception) error
	UpdateReceptionStatus(ctx context.Context, reception domain.Reception) error
	CreateProduct(ctx context.Context, product domain.Product) error
	DeleteLastProduct(ctx context.Context, receptionID uuid.UUID) error
	GetPvzList(ctx context.Context, filter domain.PvzFilter) ([]domain.Pvz, error)
	GetReceptionsByPvz(ctx context.Context, pvzID uuid.UUID, filter domain.PvzFilter) ([]domain.Reception, error)
	GetProductsByReception(ctx context.Context, receptionID uuid.UUID) ([]domain.Product, error)
}
