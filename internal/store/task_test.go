package store

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dragon-huang0403/todo-go/internal/db"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()
		expectedTask := &models.Task{
			ID:        taskID,
			Name:      gofakeit.Name(),
			Status:    models.TaskStatus(gofakeit.Number(0, 1)),
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}

		// stubs
		m.mockDB.EXPECT().Get(db.Task, taskID).Return(interface{}(expectedTask), nil)

		// assert
		task, err := m.store.GetTask(taskID)
		require.NoError(t, err)
		require.Equal(t, expectedTask, task)
	})

	t.Run("not found", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()

		// stubs
		m.mockDB.EXPECT().Get(db.Task, taskID).Return(nil, db.ErrNotFound)

		// assert
		task, err := m.store.GetTask(taskID)
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, task)
	})
}

func TestListTasks(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskCount := gofakeit.Number(1, 10)
		expectedTasks := make([]*models.Task, 0, taskCount)
		mockReturned := make([]interface{}, 0, taskCount)
		for i := 0; i < taskCount; i++ {
			task := &models.Task{
				ID:        uuid.New(),
				Name:      gofakeit.Name(),
				Status:    models.TaskStatus(gofakeit.Number(0, 1)),
				CreatedAt: gofakeit.Date(),
				UpdatedAt: gofakeit.Date(),
			}
			expectedTasks = append(expectedTasks, task)
			mockReturned = append(mockReturned, interface{}(task))
		}

		// stubs
		m.mockDB.EXPECT().List(db.Task).Return(mockReturned, nil)

		// assert
		tasks, err := m.store.ListTasks()
		require.NoError(t, err)
		require.Equal(t, expectedTasks, tasks)
	})

	t.Run("no rows", func(t *testing.T) {
		m := setup(t)

		// stubs
		m.mockDB.EXPECT().List(db.Task).Return([]interface{}{}, nil)

		// assert
		tasks, err := m.store.ListTasks()
		require.NoError(t, err)
		require.NotNil(t, tasks)
		require.Empty(t, tasks)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		arg := CreateTaskParams{
			Name:   gofakeit.Name(),
			Status: models.TaskStatus(gofakeit.Number(0, 1)),
		}

		// stubs
		m.mockDB.EXPECT().Create(db.Task, gomock.Any(), gomock.Any()).Return(nil)

		// assert
		task, err := m.store.CreateTask(arg)
		require.NoError(t, err)
		require.NotNil(t, task)

		require.NotZero(t, task.ID)
		require.Equal(t, arg.Name, task.Name)
		require.Equal(t, arg.Status, task.Status)
		require.WithinDuration(t, time.Now(), task.CreatedAt, time.Second)
		require.WithinDuration(t, time.Now(), task.UpdatedAt, time.Second)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()
		arg := UpdateTaskParams{
			ID:     taskID,
			Name:   gofakeit.Name(),
			Status: models.TaskStatus(gofakeit.Number(0, 1)),
		}

		oldTask := &models.Task{
			ID:        taskID,
			Name:      gofakeit.Name(),
			Status:    models.TaskStatus(gofakeit.Number(0, 1)),
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}

		// stubs
		m.mockDB.EXPECT().Get(db.Task, taskID).Return(oldTask, nil)
		m.mockDB.EXPECT().Update(db.Task, taskID, gomock.Any()).Return(nil)

		// assert
		task, err := m.store.UpdateTask(arg)
		require.NoError(t, err)
		require.NotNil(t, task)

		require.Equal(t, taskID, task.ID)
		require.Equal(t, arg.Name, task.Name)
		require.Equal(t, arg.Status, task.Status)
		require.WithinDuration(t, time.Now(), task.UpdatedAt, time.Second)
		require.Equal(t, oldTask.CreatedAt, task.CreatedAt)
	})

	t.Run("not found", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()
		arg := UpdateTaskParams{
			ID:     taskID,
			Name:   gofakeit.Name(),
			Status: models.TaskStatus(gofakeit.Number(0, 1)),
		}

		// stubs
		m.mockDB.EXPECT().Get(db.Task, taskID).Return(nil, db.ErrNotFound)

		// assert
		task, err := m.store.UpdateTask(arg)
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, task)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()

		// stubs
		m.mockDB.EXPECT().Delete(db.Task, taskID).Return(nil)

		// assert
		err := m.store.DeleteTask(taskID)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		m := setup(t)

		// prepare
		taskID := uuid.New()

		// stubs
		m.mockDB.EXPECT().Delete(db.Task, taskID).Return(db.ErrNotFound)

		// assert
		err := m.store.DeleteTask(taskID)
		require.ErrorIs(t, err, ErrNotFound)
	})
}
