package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (r *repository) GetProductsByReception(ctx context.Context, receptionID uuid.UUID) ([]domain.Product, error) {
	query, args, err := squirrel.
		Select("id", "reception_id", "type", "created_at").
		From("products").
		Where(squirrel.Eq{"reception_id": receptionID}).
		OrderBy("created_at ASC").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build product query: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute product query: %w", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Reception.ID, &p.Type, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
