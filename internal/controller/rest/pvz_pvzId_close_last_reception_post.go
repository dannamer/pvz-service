package rest

import (
	"context"
	"errors"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) PvzPvzIdCloseLastReceptionPost(ctx context.Context, params api.PvzPvzIdCloseLastReceptionPostParams) (api.PvzPvzIdCloseLastReceptionPostRes, error) {
	claims, _ := domain.ClaimsFromContext(ctx)

	res, err := h.use.CloseLastReception(ctx, domain.Reception{
		PVZID: params.PvzId,
		CreatedBy: domain.User{
			ID:   claims.UserID,
			Role: claims.Role,
		},
	})
	if err != nil {
		if errors.Is(err, domain.ErrReceptionAlreadyClosed) {
			return &api.PvzPvzIdCloseLastReceptionPostBadRequest{
				Message: domain.ErrReceptionAlreadyClosed.Error(),
			}, nil
		}
		h.log.With(ctx).Error(ctx, "PvzPvzIdCloseLastReceptionPost", err)
		return nil, domain.ErrInternalServerError
	}

	return &api.Reception{
		ID:       api.NewOptUUID(res.ID),
		DateTime: res.UpdatedAt,
		PvzId:    res.PVZID,
		Status:   api.ReceptionStatus(res.Status),
	}, nil
}
