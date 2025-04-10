package jwt

import (
	"time"

	jwt_std "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (j *jwt) GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt_std.NewWithClaims(jwt_std.SigningMethodHS256, jwt_std.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(j.key))
}
