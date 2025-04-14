package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	ReceptionStatusInProgress = "in_progress"
	ReceptionStatusClosed     = "closed"
)

type Reception struct {
	ID        uuid.UUID
	PVZID     uuid.UUID
	CreatedBy User
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
