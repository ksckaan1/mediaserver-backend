package repository

import (
	"mediaserver/internal/infrastructure/repository/sqlcgen"
)

type Repository struct {
	queries *sqlcgen.Queries
}

func New(db sqlcgen.DBTX) (*Repository, error) {
	return &Repository{
		queries: sqlcgen.New(db),
	}, nil
}
