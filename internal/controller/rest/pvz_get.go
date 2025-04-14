package rest

import (
	"context"

	"github.com/dannamer/pvz-service/internal/domain"
	"github.com/dannamer/pvz-service/internal/generated/api"
)

func (h *handler) PvzGet(ctx context.Context, params api.PvzGetParams) ([]api.PvzGetOKItem, error) {
	filter := domain.PvzFilter{
		Page:  1,
		Limit: 10,
	}
	if params.StartDate.Set {
		filter.StartDate = &params.StartDate.Value
	}
	if params.EndDate.Set {
		filter.EndDate = &params.EndDate.Value
	}
	if params.Page.Set {
		filter.Page = params.Page.Value
	}
	if params.Limit.Set {
		filter.Limit = params.Limit.Value
	}
	pvzs, err := h.use.GetPvzList(ctx, filter)
	if err != nil {
		h.log.With(ctx).Error(ctx, "PvzGet", err)
		return nil, domain.ErrInternalServerError
	}
	result := make([]api.PvzGetOKItem, 0, len(pvzs))
	for _, pvz := range pvzs {
		pvzReceptions := make([]api.PvzGetOKItemReceptionsItem, 0, len(pvz.Receptions))
		for _, reception := range pvz.Receptions {
			pvzProduct := make([]api.Product, 0, len(reception.Products))
			for _, product := range reception.Products {
				pvzProduct = append(pvzProduct, api.Product{
					ID:          api.NewOptUUID(product.ID),
					DateTime:    api.NewOptDateTime(product.CreatedAt),
					Type:        api.ProductType(product.Type),
					ReceptionId: product.Reception.ID,
				})
			}
			pvzReceptions = append(pvzReceptions, api.PvzGetOKItemReceptionsItem{
				Reception: api.NewOptReception(api.Reception{
					ID:       api.NewOptUUID(reception.Reception.ID),
					DateTime: reception.Reception.CreatedAt,
					PvzId:    reception.Reception.PVZID,
					Status:   api.ReceptionStatus(reception.Reception.Status),
				}),
				Products: pvzProduct,
			})
		}
		result = append(result, api.PvzGetOKItem{
			Pvz: api.NewOptPVZ(api.PVZ{
				ID:               api.NewOptUUID(pvz.Pvz.ID),
				RegistrationDate: api.NewOptDateTime(pvz.Pvz.CreatedAt),
				City:             api.PVZCity(pvz.Pvz.City),
			}),
			Receptions: pvzReceptions,
		})
	}
	return result, nil
}
