package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerId, userID int64) error {
	query := `
		INSERT INTO followers (
			user_id, 
			follower_id
		)
		VALUES (
			$1,
			$2
		)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return err
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerId, userID int64) error {
	query := `
		DELETE FROM followers 
		WHERE
			user_id = $1 AND follower_id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerId)
	return err
}
