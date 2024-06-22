package httptest

import (
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestListTasks(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		n := gofakeit.Number(1, 100)
		tasks := m.prepareTasks(t, n)
		result := m.expect.GET("/tasks").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("data").Array()
		result.Length().IsEqual(n)

		result.Every(func(index int, value *httpexpect.Value) {
			item := value.Object()
			item.Value("id").IsEqual(tasks[index].ID)
			item.Value("name").IsEqual(tasks[index].Name)
			item.Value("status").IsEqual(tasks[index].Status)
			item.Value("created_at").IsEqual(tasks[index].CreatedAt)
			item.Value("updated_at").IsEqual(tasks[index].UpdatedAt)
		})
	})

	t.Run("no rows", func(t *testing.T) {
		m := setup(t)
		m.expect.GET("/tasks").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("data").Array().Length().IsEqual(0)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		name := gofakeit.Name()
		status := randomTaskStatus()

		// assert
		result := m.expect.POST("/tasks").
			WithJSON(map[string]interface{}{
				"name":   name,
				"status": status,
			}).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		result.Value("data").Object().Value("name").IsEqual(name)
		result.Value("data").Object().Value("status").IsEqual(status)
		result.Value("data").Object().Value("created_at").String().AsDateTime(time.RFC3339)
		result.Value("data").Object().Value("updated_at").String().AsDateTime(time.RFC3339)

		// check database
		rawId := result.Value("data").Object().Value("id").String().Raw()
		id, err := uuid.Parse(rawId)
		require.NoError(t, err)

		task, err := m.store.GetTask(id)
		require.NoError(t, err)

		require.Equal(t, name, task.Name)
		require.Equal(t, status, task.Status)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		task := m.prepareTask(t)
		name := gofakeit.Name()
		status := randomTaskStatus()

		// assert
		result := m.expect.PUT("/tasks/" + task.ID.String()).
			WithJSON(map[string]interface{}{
				"name":   name,
				"status": status,
			}).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		result.Value("data").Object().Value("name").IsEqual(name)
		result.Value("data").Object().Value("status").IsEqual(status)
		result.Value("data").Object().Value("created_at").String().AsDateTime(time.RFC3339)
		result.Value("data").Object().Value("updated_at").String().AsDateTime(time.RFC3339)

		// check database
		task, err := m.store.GetTask(task.ID)
		require.NoError(t, err)

		require.Equal(t, name, task.Name)
		require.Equal(t, status, task.Status)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		task := m.prepareTask(t)

		// assert
		m.expect.DELETE("/tasks/" + task.ID.String()).
			Expect().
			Status(http.StatusOK)

		// check database
		_, err := m.store.GetTask(task.ID)
		require.ErrorIs(t, err, controller.ErrNotFound)
	})
}
