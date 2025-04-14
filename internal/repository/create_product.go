package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) CreateProduct(ctx context.Context, product domain.Product) error {
	query, args, err := squirrel.
		Insert("products").
		Columns("id", "pvz_id", "reception_id", "type", "created_at").
		Values(product.ID, product.PVZID, product.Reception.ID, product.Type, product.CreatedAt).
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
