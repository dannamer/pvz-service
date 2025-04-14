package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (u *usecase) CloseLastReception(ctx context.Context, reception domain.Reception) (domain.Reception, error) {
	reception, err := u.repo.GetLastReception(ctx, reception.PVZID)
	if err != nil {
		return domain.Reception{}, fmt.Errorf("failed to get last reception: %w", err)
	}

	if reception.ID == uuid.Nil || reception.Status == domain.ReceptionStatusClosed {
		return domain.Reception{}, domain.ErrReceptionAlreadyClosed
	}

	reception.Status = domain.ReceptionStatusClosed
	reception.UpdatedAt = time.Now()

	if err := u.repo.UpdateReceptionStatus(ctx, reception); err != nil {
		return domain.Reception{}, fmt.Errorf("failed to update reception status: %w", err)
	}

	return reception, nil
}
