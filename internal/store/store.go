package store

import "github.com/dragon-huang0403/todo-go/internal/db"

var (
	ErrNotFound = db.ErrNotFound
	ErrConflict = db.ErrAlreadyExists
)

type Store struct {
	db db.Database
}

func New(database db.Database) *Store {
	return &Store{
		db: database,
	}
}
