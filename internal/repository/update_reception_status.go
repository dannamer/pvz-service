package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) UpdateReceptionStatus(ctx context.Context, reception domain.Reception) error {
	query, args, err := squirrel.
		Update("receptions").
		Set("status", reception.Status).
		Set("updated_at", reception.UpdatedAt).
		Where("id = ?", reception.ID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	return nil
}
