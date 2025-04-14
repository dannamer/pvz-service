package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (r *repository) GetReceptionsByPvz(ctx context.Context, pvzID uuid.UUID, filter domain.PvzFilter) ([]domain.Reception, error) {
	sb := squirrel.
		Select("id", "pvz_id", "status", "created_at").
		From("receptions").
		Where(squirrel.Eq{"pvz_id": pvzID}).
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("created_at DESC")

	if filter.StartDate != nil {
		sb = sb.Where(squirrel.GtOrEq{"created_at": *filter.StartDate})
	}
	if filter.EndDate != nil {
		sb = sb.Where(squirrel.LtOrEq{"created_at": *filter.EndDate})
	}

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build reception query: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute reception query: %w", err)
	}
	defer rows.Close()

	var receptions []domain.Reception
	for rows.Next() {
		var rec domain.Reception
		if err := rows.Scan(&rec.ID, &rec.PVZID, &rec.Status, &rec.CreatedAt); err != nil {
			return nil, err
		}
		receptions = append(receptions, rec)
	}

	return receptions, nil
}
