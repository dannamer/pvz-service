package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	query, args, err := squirrel.
		Select("id", "email", "password", "role").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
