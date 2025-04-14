package store

import (
	"context"
	"database/sql"

	"github.com/revacholiere-moralist/cleanarch-go/internal/model"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO posts (username, password, email)
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
