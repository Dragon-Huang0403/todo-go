package controller

import "github.com/dragon-huang0403/todo-go/internal/store"

var (
	ErrNotFound = store.ErrNotFound
)

type Controller struct {
	Task Task
}

func New(store store.Store) *Controller {
	return &Controller{
		Task: NewTask(store),
	}
}
