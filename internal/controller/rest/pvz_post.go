package rest

import (
	"context"
	"errors"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) PvzPost(ctx context.Context, req *api.PVZ) (api.PvzPostRes, error) {
	claims, _ := domain.ClaimsFromContext(ctx)
	pvz, err := h.use.RegisterPvz(ctx, domain.Pvz{
		CreatedBy: domain.User{
			ID:   claims.UserID,
			Role: claims.Role,
		},
		City: string(req.GetCity()),
	})
	if err != nil {
		if errors.Is(err, domain.ErrModeratorAccessDenied) {
			return &api.PvzPostForbidden{
				Message: domain.ErrModeratorAccessDenied.Error(),
			}, nil
		}
		h.log.With(ctx).Error(ctx, "PvzPost", err)
		return nil, domain.ErrInternalServerError
	}
	return &api.PVZ{
		ID:               api.NewOptUUID(pvz.ID),
		RegistrationDate: api.NewOptDateTime(pvz.CreatedAt),
		City:             api.PVZCity(pvz.City),
	}, nil
}
