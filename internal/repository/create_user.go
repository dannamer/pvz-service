package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) CreateUser(ctx context.Context, user domain.User) error {
	query, args, err := squirrel.
		Insert("users").
		Columns("id", "email", "password", "role", "created_at", "updated_at").
		Values(user.ID, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
