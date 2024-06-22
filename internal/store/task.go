package store

import (
	"time"

	"github.com/dragon-huang0403/todo-go/internal/db"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/google/uuid"
)

func (s *Store) GetTask(id uuid.UUID) (*models.Task, error) {
	task, err := s.db.Get(db.Task, id)
	if err != nil {
		return nil, err
	}

	return models.Task{}.FromDB(task)
}

func (s *Store) ListTasks() ([]*models.Task, error) {
	tasks, err := s.db.List(db.Task)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Task, 0, len(tasks))
	for _, task := range tasks {
		t, err := models.Task{}.FromDB(task)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

type CreateTaskParams struct {
	Name   string
	Status models.TaskStatus
}

func (s *Store) CreateTask(params CreateTaskParams) (*models.Task, error) {
	task := &models.Task{
		ID:        uuid.New(),
		Name:      params.Name,
		Status:    params.Status,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.db.Create(db.Task, task.ID, task); err != nil {
		return nil, err
	}

	return task, nil
}

type UpdateTaskParams struct {
	ID     uuid.UUID
	Name   string
	Status models.TaskStatus
}

func (s *Store) UpdateTask(params UpdateTaskParams) (*models.Task, error) {
	task, err := s.GetTask(params.ID)
	if err != nil {
		return nil, err
	}

	task.Name = params.Name
	task.Status = params.Status
	task.UpdatedAt = time.Now().UTC()

	if err := s.db.Update(db.Task, task.ID, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Store) DeleteTask(id uuid.UUID) error {
	return s.db.Delete(db.Task, id)
}
