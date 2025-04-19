package store

import (
	"context"
	"database/sql"
	"errors"

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
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

func (s *UsersStore) GetByID(ctx context.Context, userID int64) (*model.User, error) {
	var user model.User

	query := `
		SELECT 
			id,
			email,
			username,
			password
		FROM public.users
		WHERE id = ($1)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return &user, ErrNotFound
		default:
			return &user, err
		}
	}

	return &user, nil
}
