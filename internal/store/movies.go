package store

import (
	"context"
	"database/sql"
)

type MoviesStore struct {
	db *sql.DB
}

func (s *MoviesStore) Create(ctx context.Context) error {
	return nil
}
