package usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/pvz-service/internal/domain"
)

func (u *usecase) GetPvzList(ctx context.Context, pvzFilter domain.PvzFilter) ([]domain.PvzWithReceptions, error) {
	pvzs, err := u.repo.GetPvzList(ctx, pvzFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to get pvz list: %w", err)
	}

	result := make([]domain.PvzWithReceptions, 0, len(pvzs))

	for _, pvz := range pvzs {
		receptions, err := u.repo.GetReceptionsByPvz(ctx, pvz.ID, pvzFilter)
		if err != nil {
			return nil, fmt.Errorf("failed to get receptions for pvz %s: %w", pvz.ID, err)
		}

		receptionWithProducts := make([]domain.ReceptionWithProducts, 0, len(receptions))

		for _, reception := range receptions {
			products, err := u.repo.GetProductsByReception(ctx, reception.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to get products for reception %s: %w", reception.ID, err)
			}

			receptionWithProducts = append(receptionWithProducts, domain.ReceptionWithProducts{
				Reception: reception,
				Products:  products,
			})
		}

		result = append(result, domain.PvzWithReceptions{
			Pvz:        pvz,
			Receptions: receptionWithProducts,
		})
	}

	return result, nil
}
