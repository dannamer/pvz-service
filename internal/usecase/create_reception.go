package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (u *usecase) CreateReception(ctx context.Context, reception domain.Reception) (domain.Reception, error) {
	if reception.CreatedBy.Role != domain.RoleEmployee {
		return domain.Reception{}, domain.ErrEmployeeAccessDenied
	}

	lastReception, err := u.repo.GetLastReception(ctx, reception.PVZID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return domain.Reception{}, fmt.Errorf("failed to get last reception: %w", err)
	}

	if lastReception.Status == domain.ReceptionStatusInProgress {
		return domain.Reception{}, domain.ErrReceptionAlreadyExists
	}

	reception = domain.Reception{
		ID:        uuid.New(),
		PVZID:     reception.PVZID,
		CreatedBy: reception.CreatedBy,
		Status:    domain.ReceptionStatusInProgress,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = u.repo.CreateReception(ctx, reception); err != nil {
		return domain.Reception{}, fmt.Errorf("failed to create reception: %w", err)
	}

	return reception, nil
}
