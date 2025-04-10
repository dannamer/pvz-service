package rest

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
)

type usecase interface {
	RegisterDummyUserWithToken(ctx context.Context, role string) (string, error)
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (string, error)
}