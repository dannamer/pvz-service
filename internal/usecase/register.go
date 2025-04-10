package usecase

import (
	"context"
	"errors"
	"net/mail"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, user domain.User) (domain.User, error) {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return domain.User{}, domain.ErrInvalidEmailFormat
	}

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	user.Password = string(hashed)

	if err := u.repo.CreateUser(ctx, user); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, domain.ErrEmailAlreadyExists
		}

		return domain.User{}, err
	}

	return user, nil
}
