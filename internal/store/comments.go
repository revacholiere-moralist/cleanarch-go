package store

import (
	"context"
	"database/sql"

	"github.com/revacholiere-moralist/cleanarch-go/internal/model"
)

type CommentsStore struct {
	db *sql.DB
}

func (s *CommentsStore) GetByPostID(ctx context.Context, postID int64) ([]model.Comment, error) {
	query := `
		SELECT 	
			comments.id,
			comments.post_id,
			comments.user_id,
			comments.content,
			comments.created_at,
			users.username,
			users.id 
		FROM comments
		JOIN users 
			ON comments.user_id = users.id
		WHERE comments.post_id = $1
		ORDER BY comments.created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []model.Comment{}
	for rows.Next() {
		var c model.Comment
		c.User = model.User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
