package usecase

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(ctx context.Context, user domain.User) (string, error) {
	storedUser, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := u.jwt.GenerateToken(domain.Claims{
		UserID: storedUser.ID,
		Role:   storedUser.Role,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
