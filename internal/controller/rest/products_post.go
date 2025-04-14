package rest

import (
	"context"
	"errors"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) ProductsPost(ctx context.Context, req *api.ProductsPostReq) (api.ProductsPostRes, error) {
	claims, _ := domain.ClaimsFromContext(ctx)
	product, err := h.use.AddProduct(ctx, domain.Product{
		Type:  string(req.Type),
		PVZID: req.PvzId,
		Reception: domain.Reception{
			CreatedBy: domain.User{
				ID:   claims.UserID,
				Role: claims.Role,
			},
		},
	})

	if err != nil {
		if errors.Is(err, domain.ErrEmployeeAccessDenied) {
			return &api.ProductsPostForbidden{
				Message: domain.ErrEmployeeAccessDenied.Error(),
			}, nil
		}
		if errors.Is(err, domain.ErrReceptionAlreadyClosed) {
			return &api.ProductsPostBadRequest{
				Message: domain.ReceptionStatusClosed,
			}, nil
		}
		h.log.With(ctx).Error(ctx, "ProductsPost", err)
		return nil, domain.ErrInternalServerError
	}

	return &api.Product{
		ID:          api.NewOptUUID(product.ID),
		DateTime:    api.NewOptDateTime(product.CreatedAt),
		Type:        api.ProductType(product.Type),
		ReceptionId: product.Reception.ID,
	}, nil
}
