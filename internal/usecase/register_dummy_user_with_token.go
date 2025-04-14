package usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (u *usecase) RegisterDummyUserWithToken(ctx context.Context, role string) (string, error) {
	userID := uuid.New()

	if _, err := u.Register(ctx, domain.User{
		ID:       userID,
		Email:    fmt.Sprintf("dummy_%s_%s@example.com", role, userID.String()[:8]),
		Password: "password",
		Role:     role,
	}); err != nil {
		return "", err
	}

	token, err := u.jwt.GenerateToken(domain.Claims{
		UserID: userID,
		Role:   role,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
