package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) GetPvzList(ctx context.Context, filter domain.PvzFilter) ([]domain.Pvz, error) {
	sb := squirrel.
		Select("id", "city", "created_at").
		From("pvz").
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("created_at DESC").
		Offset(uint64((filter.Page - 1) * filter.Limit)).
		Limit(uint64(filter.Limit))

	if filter.StartDate != nil {
		sb = sb.Where(squirrel.GtOrEq{"created_at": *filter.StartDate})
	}
	if filter.EndDate != nil {
		sb = sb.Where(squirrel.LtOrEq{"created_at": *filter.EndDate})
	}

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build pvz list query: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute pvz list query: %w", err)
	}
	defer rows.Close()

	var pvzs []domain.Pvz
	for rows.Next() {
		var p domain.Pvz
		if err := rows.Scan(&p.ID, &p.City, &p.CreatedAt); err != nil {
			return nil, err
		}
		pvzs = append(pvzs, p)
	}

	return pvzs, nil
}
