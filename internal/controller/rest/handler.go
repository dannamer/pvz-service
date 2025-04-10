package rest

import "github.com/dannamer/pvz-service/internal/infrastructure/logger"

type handler struct {
	use usecase
	log logger.Logger
}

func NewHandlers(use usecase, log logger.Logger) *handler {
	return &handler{
		use: use,
		log: log,
	}
}
