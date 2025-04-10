package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) CreateUser(ctx context.Context, user domain.User) error {
	query, args, err := squirrel.
		Insert("users").
		Columns("id", "email", "password", "role").
		Values(user.ID, user.Email, user.Password, user.Role).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}
