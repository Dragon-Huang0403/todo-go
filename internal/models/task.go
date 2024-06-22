package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrConvertFailed = errors.New("failed to convert interface to Task")
)

type TaskStatus int

const (
	TaskStatusIncomplete TaskStatus = iota
	TaskStatusCompleted
)

type Task struct {
	ID uuid.UUID `json:"id" validate:"required" format:"uuid"`

	// task name
	Name string `json:"name" validate:"required" example:"account name"`

	// 0 represents an incomplete task, 1 represents a completed task
	Status    TaskStatus `json:"status" validate:"required" swaggertype:"integer" example:"0"`
	CreatedAt time.Time  `json:"created_at" validate:"required" format:"date-time"`
	UpdatedAt time.Time  `json:"updated_at" validate:"required" format:"date-time"`
}

func (Task) FromDB(v interface{}) (*Task, error) {
	task, ok := v.(*Task)
	if !ok {
		return nil, ErrConvertFailed
	}
	return task, nil
}
