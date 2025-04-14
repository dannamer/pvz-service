package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateReception(ctx context.Context, reception domain.Reception) error {
	query, args, err := squirrel.
		Insert("receptions").
		Columns("id", "pvz_id", "created_by", "status", "created_at", "updated_at").
		Values(reception.ID, reception.PVZID, reception.CreatedBy.ID, reception.Status, reception.CreatedAt, reception.UpdatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrPvzNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
