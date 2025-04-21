package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/revacholiere-moralist/cleanarch-go/internal/model"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = 30 * time.Second
)

type Storage struct {
	Posts interface {
		Create(context.Context, *model.Post) error
		GetByID(ctx context.Context, postID int64) (*model.Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *model.Post) error
	}

	Users interface {
		Create(context.Context, *model.User) error
		GetByID(context.Context, int64) (*model.User, error)
	}

	Comments interface {
		GetByPostID(context.Context, int64) ([]model.Comment, error)
		AddCommentsToPost(context.Context, int64, int64, string, *model.Comment) error
	}

	Followers interface {
		Follow(ctx context.Context, followerId, userID int64) error
		Unfollow(ctx context.Context, followerId, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostsStore{db},
		Users:     &UsersStore{db},
		Comments:  &CommentsStore{db},
		Followers: &FollowerStore{db},
	}
}
