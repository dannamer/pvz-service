package rest

import (
	"context"
	"errors"
	"fmt"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) LoginPost(ctx context.Context, req *api.LoginPostReq) (api.LoginPostRes, error) {
	token, err := h.use.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			h.log.With(ctx).Warn(ctx, fmt.Sprintf("login failed: invalid credentials, email=%s", req.Email))
			return &api.Error{Message: domain.ErrInvalidCredentials.Error()}, nil
		}

		h.log.With(ctx).Error(ctx, "LoginPost", err)
		return nil, domain.ErrInternalServerError
	}

	h.log.With(ctx).Info(ctx, fmt.Sprintf("login successful: email=%s", req.Email))

	tok := api.Token(token)
	return &tok, nil
}
