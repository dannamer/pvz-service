package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID
	PVZID     uuid.UUID
	Reception Reception
	Type      string
	CreatedAt time.Time
}
