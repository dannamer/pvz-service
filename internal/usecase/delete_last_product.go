package usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/pvz-service/internal/domain"
)

func (u *usecase) DeleteLastProduct(ctx context.Context, pvz domain.Pvz) error {
	if pvz.CreatedBy.Role != domain.RoleEmployee {
		return domain.ErrEmployeeAccessDenied
	}

	reception, err := u.repo.GetLastReception(ctx, pvz.ID)
	if err != nil {
		return fmt.Errorf("failed to get last reception: %w", err)
	}
	if reception.Status != domain.ReceptionStatusInProgress {
		return domain.ErrReceptionAlreadyClosed
	}

	if err = u.repo.DeleteLastProduct(ctx, reception.ID); err != nil {
		return fmt.Errorf("failed to delete last product: %w", err)
	}
	return nil
}
