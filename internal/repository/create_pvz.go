package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/pvz-service/internal/domain"
)

func (r *repository) CreatePvz(ctx context.Context, pvz domain.Pvz) error {
	query, args, err := squirrel.
		Insert("pvz").
		Columns("id", "created_by", "city", "created_at", "updated_at").
		Values(pvz.ID, pvz.CreatedBy.ID, pvz.City, pvz.CreatedAt, pvz.UpdatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}
