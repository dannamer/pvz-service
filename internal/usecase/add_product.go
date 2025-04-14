package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (u *usecase) AddProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if product.Reception.CreatedBy.Role != domain.RoleEmployee {
		return domain.Product{}, domain.ErrEmployeeAccessDenied
	}

	reception, err := u.repo.GetLastReception(ctx, product.PVZID)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to get last reception: %w", err)
	}
	if reception.Status != domain.ReceptionStatusInProgress {
		return domain.Product{}, domain.ErrReceptionAlreadyClosed
	}
	product.ID = uuid.New()
	product.CreatedAt = time.Now()
	product.Reception.ID = reception.ID
	err = u.repo.CreateProduct(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to add product: %w", err)
	}

	return product, nil
}
