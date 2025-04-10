package jwt

import (
	"errors"
	"fmt"

	jwt_std "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *jwt) ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt_std.Parse(
		tokenString,
		func(t *jwt_std.Token) (interface{}, error) {
			if t.Method != jwt_std.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(j.key), nil
		},
	)

	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(jwt_std.MapClaims)
	if !ok || !token.Valid {
		return uuid.UUID{}, errors.New("invalid token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("sub not found in token")
	}
	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid user ID format")
	}

	return userID, nil
}
