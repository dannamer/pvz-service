package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetLastReception(ctx context.Context, pvzID uuid.UUID) (domain.Reception, error) {
	query, args, err := squirrel.
		Select("id", "pvz_id", "created_by", "status", "created_at", "updated_at").
		From("receptions").
		Where("pvz_id = ?", pvzID).
		OrderBy("created_at DESC").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return domain.Reception{}, fmt.Errorf("failed to build query: %w", err)
	}

	var reception domain.Reception
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&reception.ID,
		&reception.PVZID,
		&reception.CreatedBy.ID,
		&reception.Status,
		&reception.CreatedAt,
		&reception.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return domain.Reception{}, nil
	}

	if err != nil {
		return domain.Reception{}, fmt.Errorf("failed to scan reception: %w", err)
	}

	return reception, nil
}
