package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/revacholiere-moralist/cleanarch-go/internal/model"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *model.Post) error
		GetByID(ctx context.Context, postID int64) (*model.Post, error)
	}

	Users interface {
		Create(context.Context, *model.User) error
	}

	Comments interface {
		GetByPostID(context.Context, int64) ([]model.Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostsStore{db},
		Users:    &UsersStore{db},
		Comments: &CommentsStore{db},
	}
}
