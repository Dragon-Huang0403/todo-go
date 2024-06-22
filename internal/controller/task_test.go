package controller

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/dragon-huang0403/todo-go/internal/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		arg := CreateTaskParams{
			Name:   gofakeit.Name(),
			Status: models.TaskStatus(gofakeit.Number(0, 1)),
		}

		expectedTask := models.Task{
			ID:        uuid.New(),
			Name:      arg.Name,
			Status:    arg.Status,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// stubs
		m.mockStore.EXPECT().CreateTask(store.CreateTaskParams(arg)).Return(&expectedTask, nil)

		// assert
		task, err := m.controller.Task.Create(ctx, arg)
		require.NoError(t, err)
		require.NotNil(t, task)

		require.Equal(t, expectedTask.ID, task.ID)
		require.Equal(t, arg.Name, task.Name)
		require.Equal(t, arg.Status, task.Status)
		require.WithinDuration(t, expectedTask.CreatedAt, task.CreatedAt, time.Second)
		require.WithinDuration(t, expectedTask.UpdatedAt, task.UpdatedAt, time.Second)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		id := uuid.New()

		// stubs
		m.mockStore.EXPECT().DeleteTask(id).Return(nil)

		// assert
		err := m.controller.Task.Delete(ctx, id)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		id := uuid.New()

		// stubs
		m.mockStore.EXPECT().DeleteTask(id).Return(store.ErrNotFound)

		// assert
		err := m.controller.Task.Delete(ctx, id)
		require.ErrorIs(t, err, ErrNotFound)
	})
}

func TestGetTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		id := uuid.New()

		expectedTask := models.Task{
			ID:        id,
			Name:      gofakeit.Name(),
			Status:    models.TaskStatus(gofakeit.Number(0, 1)),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// stubs
		m.mockStore.EXPECT().GetTask(id).Return(&expectedTask, nil)

		// assert
		task, err := m.controller.Task.Get(ctx, id)
		require.NoError(t, err)
		require.NotNil(t, task)

		require.Equal(t, expectedTask.ID, task.ID)
		require.Equal(t, expectedTask.Name, task.Name)
		require.Equal(t, expectedTask.Status, task.Status)
		require.WithinDuration(t, expectedTask.CreatedAt, task.CreatedAt, time.Second)
		require.WithinDuration(t, expectedTask.UpdatedAt, task.UpdatedAt, time.Second)
	})

	t.Run("not found", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		id := uuid.New()

		// stubs
		m.mockStore.EXPECT().GetTask(id).Return(nil, store.ErrNotFound)

		// assert
		task, err := m.controller.Task.Get(ctx, id)
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, task)
	})
}

func TestListTasks(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		n := gofakeit.Number(1, 10)

		// arrange
		expectedTasks := make([]*models.Task, 0)
		for range n {
			task := models.Task{
				ID:        uuid.New(),
				Name:      gofakeit.Name(),
				Status:    models.TaskStatus(gofakeit.Number(0, 1)),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			expectedTasks = append(expectedTasks, &task)
		}

		// stubs
		m.mockStore.EXPECT().ListTasks().Return(expectedTasks, nil)

		// assert
		tasks, err := m.controller.Task.List(ctx)
		require.NoError(t, err)
		require.NotNil(t, tasks)
		require.Len(t, tasks, len(expectedTasks))

		for i, expectedTask := range expectedTasks {
			task := tasks[i]
			require.Equal(t, expectedTask.ID, task.ID)
			require.Equal(t, expectedTask.Name, task.Name)
			require.Equal(t, expectedTask.Status, task.Status)
			require.WithinDuration(t, expectedTask.CreatedAt, task.CreatedAt, time.Second)
			require.WithinDuration(t, expectedTask.UpdatedAt, task.UpdatedAt, time.Second)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		arg := UpdateTaskParams{
			ID:   uuid.New(),
			Name: gofakeit.Name(),
		}

		expectedTask := models.Task{
			ID:        arg.ID,
			Name:      arg.Name,
			Status:    models.TaskStatus(gofakeit.Number(0, 1)),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// stubs
		m.mockStore.EXPECT().UpdateTask(store.UpdateTaskParams(arg)).Return(&expectedTask, nil)

		// assert
		task, err := m.controller.Task.Update(ctx, arg)
		require.NoError(t, err)
		require.NotNil(t, task)

		require.Equal(t, expectedTask.ID, task.ID)
		require.Equal(t, arg.Name, task.Name)
		require.Equal(t, expectedTask.Status, task.Status)
		require.WithinDuration(t, expectedTask.CreatedAt, task.CreatedAt, time.Second)
		require.WithinDuration(t, expectedTask.UpdatedAt, task.UpdatedAt, time.Second)
	})

	t.Run("not found", func(t *testing.T) {
		ctx := context.Background()
		m := setup(t)

		// arrange
		arg := UpdateTaskParams{
			ID:   uuid.New(),
			Name: gofakeit.Name(),
		}

		// stubs
		m.mockStore.EXPECT().UpdateTask(store.UpdateTaskParams(arg)).Return(nil, store.ErrNotFound)

		// assert
		task, err := m.controller.Task.Update(ctx, arg)
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, task)
	})
}
