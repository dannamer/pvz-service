package rest

import (
	"context"
	"errors"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) ReceptionsPost(ctx context.Context, req *api.ReceptionsPostReq) (api.ReceptionsPostRes, error) {
	claims, _ := domain.ClaimsFromContext(ctx)
	res, err := h.use.CreateReception(ctx, domain.Reception{
		PVZID: req.GetPvzId(),
		CreatedBy: domain.User{
			ID:   claims.UserID,
			Role: claims.Role,
		},
	})
	if err != nil {
		if errors.Is(err, domain.ErrEmployeeAccessDenied) {
			return &api.ReceptionsPostForbidden{
				Message: domain.ErrEmployeeAccessDenied.Error(),
			}, nil
		}
		if errors.Is(err, domain.ErrReceptionAlreadyExists) {
			return &api.ReceptionsPostBadRequest{
				Message: domain.ErrReceptionAlreadyExists.Error(),
			}, nil
		}
		if errors.Is(err, domain.ErrPvzNotFound) {
			return &api.ReceptionsPostBadRequest{
				Message: domain.ErrPvzNotFound.Error(),
			}, nil
		}
		h.log.With(ctx).Error(ctx, "ReceptionsPost", err)
		return nil, domain.ErrInternalServerError
	}

	return &api.Reception{
		ID:       api.NewOptUUID(res.ID),
		DateTime: res.CreatedAt,
		PvzId:    res.PVZID,
		Status:   api.ReceptionStatus(res.Status),
	}, nil
}
