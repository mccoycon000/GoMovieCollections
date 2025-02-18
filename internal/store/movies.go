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
	UserID      int64    `json:"user_id"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"releasedate"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type MovieStore struct {
	db *sql.DB
}

func (s *MovieStore) Create(ctx context.Context, movie *Movie) error {
	query := `
		INSERT INTO movies (title, description, releasedate, user_id, tags)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		movie.Title,
		movie.Description,
		movie.ReleaseDate,
		movie.UserID,
		pq.Array(movie.Tags),
	).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
