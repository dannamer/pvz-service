package jwt

import (
	"time"

	"github.com/dannamer/pvz-service/internal/domain"
	jwt_std "github.com/golang-jwt/jwt/v5"
)

func (j *jwt) GenerateToken(claims domain.Claims) (string, error) {
	claims.RegisteredClaims = jwt_std.RegisteredClaims{
		ExpiresAt: jwt_std.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt_std.NewNumericDate(time.Now()),
	}

	token := jwt_std.NewWithClaims(jwt_std.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.key))
}
