package middleware

import "github.com/dannamer/pvz-service/internal/domain"

type jwt interface {
	ParseToken(tokenStr string) (*domain.Claims, error)
}
