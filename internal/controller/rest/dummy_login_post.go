package rest

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) DummyLoginPost(ctx context.Context, req *api.DummyLoginPostReq) (api.DummyLoginPostRes, error) {
	token, err := h.use.RegisterDummyUserWithToken(ctx, string(req.Role))
	if err != nil {
		h.log.With(ctx).Error(ctx, "не удалось сгенерировать токен", err)
		return &api.Error{
			Message: domain.ErrInternalServerError.Error(),
		}, nil
	}
	h.log.With(ctx).Info(ctx, "токен успешно сгенерирован")
	tok := api.Token(token)
	return &tok, nil
}
