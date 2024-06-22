package store

import (
	"github.com/dragon-huang0403/todo-go/internal/db"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/google/uuid"
)

var (
	ErrNotFound = db.ErrNotFound
	ErrConflict = db.ErrAlreadyExists
)

type Store interface {
	GetTask(uuid.UUID) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	CreateTask(CreateTaskParams) (*models.Task, error)
	UpdateTask(UpdateTaskParams) (*models.Task, error)
	DeleteTask(uuid.UUID) error
}

type storeImpl struct {
	db db.Database
}

func New(database db.Database) Store {
	return &storeImpl{
		db: database,
	}
}
