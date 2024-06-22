package httptest

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/internal/db"
	httpserver "github.com/dragon-huang0403/todo-go/internal/http/server"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/dragon-huang0403/todo-go/internal/store"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/require"
)

type testMain struct {
	expect *httpexpect.Expect

	store store.Store
}

func setup(t *testing.T) *testMain {
	ctx := context.Background()
	db := db.New()
	store := store.New(db)
	controller := controller.New(store)
	validator := validator.New()

	server := httptest.NewServer(httpserver.NewServer(ctx, controller, validator))
	t.Cleanup(server.Close)

	expect := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return &testMain{
		expect: expect,
		store:  store,
	}
}

func (m *testMain) prepareTask(t *testing.T) *models.Task {
	task, err := m.store.CreateTask(store.CreateTaskParams{
		Name:   gofakeit.Name(),
		Status: randomTaskStatus(),
	})

	require.NoError(t, err)
	return task
}

func (m *testMain) prepareTasks(t *testing.T, count int) []*models.Task {
	tasks := make([]*models.Task, count)
	for i := 0; i < count; i++ {
		tasks[i] = m.prepareTask(t)
	}

	return tasks
}

func randomTaskStatus() models.TaskStatus {
	return models.TaskStatus(gofakeit.Number(0, 1))
}
