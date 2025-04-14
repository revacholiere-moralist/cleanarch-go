package store

import (
	"context"
	"database/sql"
	"github.com/revacholiere-moralist/cleanarch-go/internal/model"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *model.Post) error
	}

	Users interface {
		Create(context.Context, *model.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
