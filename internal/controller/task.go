package controller

import (
	"context"

	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/dragon-huang0403/todo-go/internal/store"
	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Task interface {
	Create(context.Context, CreateTaskParams) (*models.Task, error)
	Delete(context.Context, uuid.UUID) error
	Get(context.Context, uuid.UUID) (*models.Task, error)
	List(context.Context) ([]*models.Task, error)
	Update(context.Context, UpdateTaskParams) (*models.Task, error)
}

type taskImpl struct {
	store store.Store
}

func NewTask(store store.Store) Task {
	return &taskImpl{
		store: store,
	}
}

type CreateTaskParams struct {
	Name   string
	Status models.TaskStatus
}

func (t *taskImpl) Create(ctx context.Context, params CreateTaskParams) (*models.Task, error) {
	logger.Debug(ctx, "Create task", zap.Any("params", params))

	task, err := t.store.CreateTask(store.CreateTaskParams(params))
	if err != nil {
		logger.Error(ctx, "Failed to create task", zap.Error(err))
		return nil, err
	}

	return task, nil
}

func (t *taskImpl) Delete(ctx context.Context, id uuid.UUID) error {
	logger.Debug(ctx, "Delete task", zap.Any("id", id))

	if err := t.store.DeleteTask(id); err != nil {
		logger.Error(ctx, "Failed to delete task", zap.Error(err))
		return err
	}

	return nil
}

func (t *taskImpl) Get(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	logger.Debug(ctx, "Get task", zap.Any("id", id))

	task, err := t.store.GetTask(id)
	if err != nil {
		logger.Error(ctx, "Failed to get task", zap.Error(err))
		return nil, err
	}

	return task, nil
}

func (t *taskImpl) List(ctx context.Context) ([]*models.Task, error) {
	logger.Debug(ctx, "List tasks")

	tasks, err := t.store.ListTasks()
	if err != nil {
		logger.Error(ctx, "Failed to list tasks", zap.Error(err))
		return nil, err
	}

	return tasks, nil
}

type UpdateTaskParams struct {
	ID     uuid.UUID
	Name   string
	Status models.TaskStatus
}

func (t *taskImpl) Update(ctx context.Context, params UpdateTaskParams) (*models.Task, error) {
	logger.Debug(ctx, "Update task", zap.Any("params", params))

	task, err := t.store.UpdateTask(store.UpdateTaskParams(params))
	if err != nil {
		logger.Error(ctx, "Failed to update task", zap.Error(err))
		return nil, err
	}

	return task, nil
}
