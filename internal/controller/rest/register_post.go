package rest

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) RegisterPost(ctx context.Context, req *api.RegisterPostReq) (api.RegisterPostRes, error) {
	user, err := h.use.Register(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     string(req.Role),
	})
	if err != nil {
		if err == domain.ErrEmailAlreadyExists {
			return &api.Error{Message: domain.ErrEmailAlreadyExists.Error()}, nil
		}
		h.log.With(ctx).Error(ctx, "ошибка регистрации", err)
		return &api.Error{Message: domain.ErrInternalServerError.Error()}, nil
	}
	h.log.With(ctx).Info(ctx, "регистрация прошла успешна")
	return &api.User{
		ID:    api.NewOptUUID(user.ID),
		Email: user.Email,
		Role:  api.UserRole(user.Role),
	}, nil
}
