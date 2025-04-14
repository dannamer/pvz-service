package usecase

import (
	"context"
	"time"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/google/uuid"
)

func (u *usecase) RegisterPvz(ctx context.Context, pvz domain.Pvz) (domain.Pvz, error) {
	if pvz.CreatedBy.Role != domain.RoleModerator {
		return domain.Pvz{}, domain.ErrModeratorAccessDenied
	}
	pvz.ID = uuid.New()
	pvz.CreatedAt = time.Now()
	pvz.UpdatedAt = time.Now()
	if err := u.repo.CreatePvz(ctx, pvz); err != nil {
		return domain.Pvz{}, err
	}
	return pvz, nil
}
