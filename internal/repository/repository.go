package repository

import "github.com/dannamer/pvz-service/internal/infrastructure/postgres"

type repository struct {
	db postgres.PgxPool
	
}

func New(db postgres.PgxPool) *repository {
	return &repository{
		db: db,
	}
}
