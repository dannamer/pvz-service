package domain

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ctxKey string

const ClaimsKey ctxKey = "claims"

type Claims struct {
	UserID uuid.UUID `json:"sub"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

func ClaimsFromContext(ctx context.Context) (Claims, error) {
	claims, ok := ctx.Value(ClaimsKey).(Claims)
	if !ok {
		return Claims{}, errors.New("claims not found in context")
	}
	return claims, nil
}
