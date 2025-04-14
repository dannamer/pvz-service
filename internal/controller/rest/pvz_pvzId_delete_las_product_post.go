package rest

import (
	"context"
	"errors"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) PvzPvzIdDeleteLastProductPost(ctx context.Context, params api.PvzPvzIdDeleteLastProductPostParams) (api.PvzPvzIdDeleteLastProductPostRes, error) {
	claims, _ := domain.ClaimsFromContext(ctx)
	if err := h.use.DeleteLastProduct(ctx, domain.Pvz{
		ID: params.PvzId,
		CreatedBy: domain.User{
			ID:   claims.UserID,
			Role: claims.Role,
		},
	}); err != nil {
		if errors.Is(err, domain.ErrEmployeeAccessDenied) {
			return &api.PvzPvzIdDeleteLastProductPostForbidden{
				Message: domain.ErrEmployeeAccessDenied.Error(),
			}, nil
		}
		if errors.Is(err, domain.ErrReceptionAlreadyClosed) {
			return &api.PvzPvzIdDeleteLastProductPostBadRequest{
				Message: domain.ErrReceptionAlreadyClosed.Error(),
			}, nil
		}
		h.log.With(ctx).Error(ctx, "PvzPvzIdDeleteLastProductPost", err)
		return nil, domain.ErrInternalServerError
	}

	return &api.PvzPvzIdDeleteLastProductPostOK{}, nil
}
