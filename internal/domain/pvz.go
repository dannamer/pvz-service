package domain

import (
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	ID        uuid.UUID
	CreatedBy User
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PvzFilter struct {
	StartDate *time.Time
	EndDate   *time.Time
	Page      int
	Limit     int
}

type PvzWithReceptions struct {
	Pvz        Pvz
	Receptions []ReceptionWithProducts
}

type ReceptionWithProducts struct {
	Reception Reception
	Products  []Product
}
