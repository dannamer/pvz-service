package middleware

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

type middleware struct {
	jwt jwt
}

func New(jwt jwt) *middleware {
	return &middleware{
		jwt: jwt,
	}
}

func (h *middleware) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
	claims, err := h.jwt.ParseToken(t.GetToken())
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, domain.ClaimsKey, *claims)

	return ctx, nil
}
