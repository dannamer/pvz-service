package jwt

import (
	"errors"
	"fmt"

	"github.com/dannamer/pvz-service/internal/domain"
	jwt_std "github.com/golang-jwt/jwt/v5"
)

func (j *jwt) ParseToken(tokenStr string) (*domain.Claims, error) {
	token, err := jwt_std.ParseWithClaims(tokenStr, &domain.Claims{}, func(t *jwt_std.Token) (interface{}, error) {
		if t.Method != jwt_std.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
