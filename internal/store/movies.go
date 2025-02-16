package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Movie Model
type Movie struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"releasedate"`
	Tags        []string `json:"tags"`
}

type MovieStore struct {
	db *sql.DB
}

func (s *MovieStore) Create(ctx context.Context, movie *Movie) error {
	query := `
		INSERT INTO movies (title, description, releasedate, tags)
		VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.db.QueryRowContext(
		ctx, 
		query, 
		movie.Title, 
		movie.Description, 
		movie.ReleaseDate, 
		pq.Array(movie.Tags),
	).Scan(
		&movie.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
