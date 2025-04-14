package rest

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
)

type usecase interface {
	RegisterDummyUserWithToken(ctx context.Context, role string) (string, error)
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (string, error)
	RegisterPvz(ctx context.Context, pvz domain.Pvz) (domain.Pvz, error)
	CreateReception(ctx context.Context, reception domain.Reception) (domain.Reception, error)
	CloseLastReception(ctx context.Context, reception domain.Reception) (domain.Reception, error)
	AddProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	DeleteLastProduct(ctx context.Context, pvz domain.Pvz) error
	GetPvzList(ctx context.Context, pvzFilter domain.PvzFilter) ([]domain.PvzWithReceptions, error)
}
