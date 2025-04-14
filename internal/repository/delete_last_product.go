package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repository) DeleteLastProduct(ctx context.Context, receptionID uuid.UUID) error {
	query, args, err := squirrel.
		Delete("products").
		Where(`
			id = (
				WITH last_product AS (
					SELECT id
					FROM products
					WHERE reception_id = ?
					ORDER BY created_at DESC
					LIMIT 1
				)
				SELECT id FROM last_product
			)`, receptionID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	return nil
}
