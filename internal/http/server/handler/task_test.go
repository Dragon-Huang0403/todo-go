package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestListTasks(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		// prepare
		c, rec := m.prepareContext(nil)

		n := gofakeit.Number(0, 10)
		data := []*models.Task{}
		for range n {
			item := models.Task{}
			err := gofakeit.Struct(&item)
			require.NoError(t, err)
			data = append(data, &item)
		}

		// stubs
		m.mockTaskCtl.EXPECT().List(gomock.Any()).Return(data, nil)

		// assert
		err := m.handler.ListTasks()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		contentType := rec.Header().Get(echo.HeaderContentType)
		require.Equal(t, echo.MIMEApplicationJSON, contentType)

		expectedData, err := json.Marshal(data)
		require.NoError(t, err)

		expectedBody := fmt.Sprintf(`{"data":%s}`, string(expectedData))
		require.JSONEq(t, expectedBody, rec.Body.String())
	})

	t.Run("error", func(t *testing.T) {
		m := setup(t)
		// prepare
		c, rec := m.prepareContext(nil)

		// stubs
		err := gofakeit.Error()
		m.mockTaskCtl.EXPECT().List(gomock.Any()).Return(nil, err)

		// assert
		err = m.handler.ListTasks()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		payload := fmt.Sprintf(`{"name":"%s","status":%d}`, gofakeit.Name(), gofakeit.Number(0, 1))
		c, rec := m.prepareContext(strings.NewReader(payload))

		task := models.Task{}
		err := gofakeit.Struct(&task)
		require.NoError(t, err)

		// stubs
		m.mockTaskCtl.EXPECT().Create(gomock.Any(), EqJSON(t, payload)).Return(&task, nil)

		// assert
		err = m.handler.CreateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		contentType := rec.Header().Get(echo.HeaderContentType)
		require.Equal(t, echo.MIMEApplicationJSON, contentType)

		expectedData, err := json.Marshal(task)
		require.NoError(t, err)

		expectedBody := fmt.Sprintf(`{"data":%s}`, string(expectedData))
		require.JSONEq(t, string(expectedBody), rec.Body.String())
	})

	t.Run("default status value", func(t *testing.T) {
		m := setup(t)

		// prepare
		name := gofakeit.Name()
		payload := fmt.Sprintf(`{"name":"%s"}`, name)
		c, rec := m.prepareContext(strings.NewReader(payload))

		task := models.Task{}
		err := gofakeit.Struct(&task)
		require.NoError(t, err)

		// stubs
		createParams := fmt.Sprintf(`{"name":"%s","status":0}`, name)
		m.mockTaskCtl.EXPECT().Create(gomock.Any(), EqJSON(t, createParams)).Return(&task, nil)

		// assert
		err = m.handler.CreateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("bad request", func(t *testing.T) {
		testCases := []struct {
			name        string
			payload     string
			errContains string
		}{{
			name:        "empty name",
			payload:     `{"name":""}`,
			errContains: `'request.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
		}, {
			name:        "invalid json",
			payload:     `{"name":}`,
			errContains: "invalid character",
		}, {
			name:        "invalid status",
			payload:     fmt.Sprintf(`{"name":"%s","status":2}`, gofakeit.Name()),
			errContains: `'request.Status' Error:Field validation for 'Status' failed on the 'oneof' tag`,
		}}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				m := setup(t)

				// prepare
				c, rec := m.prepareContext(strings.NewReader(tc.payload))

				// assert
				err := m.handler.CreateTask()(c)
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, rec.Code)
				require.Contains(t, rec.Body.String(), tc.errContains)
			})
		}
	})

	t.Run("error", func(t *testing.T) {
		m := setup(t)

		// prepare
		payload := fmt.Sprintf(`{"name":"%s","status":%d}`, gofakeit.Name(), gofakeit.Number(0, 1))
		c, rec := m.prepareContext(strings.NewReader(payload))

		// stubs
		err := gofakeit.Error()
		m.mockTaskCtl.EXPECT().Create(gomock.Any(), EqJSON(t, payload)).Return(nil, err)

		// assert
		err = m.handler.CreateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.NewString()
		name := gofakeit.Name()
		status := gofakeit.Number(0, 1)
		payload := fmt.Sprintf(`{"name":"%s","status":%d}`, name, status)
		c, rec := m.prepareContext(strings.NewReader(payload))
		c.SetParamNames("taskId")
		c.SetParamValues(id)

		task := models.Task{}
		err := gofakeit.Struct(&task)
		require.NoError(t, err)

		// stubs
		updateParams := fmt.Sprintf(`{"name":"%s","status":%d,"id":"%s"}`, name, status, id)
		m.mockTaskCtl.EXPECT().Update(gomock.Any(), EqJSON(t, updateParams)).Return(&task, nil)

		// assert
		err = m.handler.UpdateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		contentType := rec.Header().Get(echo.HeaderContentType)
		require.Equal(t, echo.MIMEApplicationJSON, contentType)

		expectedData, err := json.Marshal(task)
		require.NoError(t, err)

		expectedBody := fmt.Sprintf(`{"data":%s}`, string(expectedData))
		require.JSONEq(t, string(expectedBody), rec.Body.String())
	})

	t.Run("bad request", func(t *testing.T) {
		testCases := []struct {
			name        string
			id          string
			payload     string
			errContains string
		}{{
			name:        "invalid id",
			id:          "invalid",
			errContains: `invalid task id`,
		}, {
			name:        "empty name",
			id:          uuid.NewString(),
			payload:     `{"name":"","status":0}`,
			errContains: `'request.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
		}, {
			name:        "empty status",
			id:          uuid.NewString(),
			payload:     `{"name":"test"}`,
			errContains: `'request.Status' Error:Field validation for 'Status' failed on the 'required' tag`,
		}, {
			name:        "invalid json",
			id:          uuid.NewString(),
			payload:     `{"name":}`,
			errContains: "invalid character",
		}, {
			name:        "invalid status",
			id:          uuid.NewString(),
			payload:     fmt.Sprintf(`{"name":"%s","status":2}`, gofakeit.Name()),
			errContains: `'request.Status' Error:Field validation for 'Status' failed on the 'oneof' tag`,
		}}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				m := setup(t)

				// prepare
				c, rec := m.prepareContext(strings.NewReader(tc.payload))
				c.SetParamNames("taskId")
				c.SetParamValues(tc.id)

				// assert
				err := m.handler.UpdateTask()(c)
				require.NoError(t, err)
				require.Equal(t, http.StatusBadRequest, rec.Code)
				require.Contains(t, rec.Body.String(), tc.errContains)
			})
		}
	})

	t.Run("not found", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.NewString()
		name := gofakeit.Name()
		status := gofakeit.Number(0, 1)
		payload := fmt.Sprintf(`{"name":"%s","status":%d}`, name, status)
		c, rec := m.prepareContext(strings.NewReader(payload))
		c.SetParamNames("taskId")
		c.SetParamValues(id)

		// stubs
		updateParams := fmt.Sprintf(`{"name":"%s","status":%d,"id":"%s"}`, name, status, id)
		m.mockTaskCtl.EXPECT().Update(gomock.Any(), EqJSON(t, updateParams)).Return(nil, controller.ErrNotFound)

		// assert
		err := m.handler.UpdateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("error", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.NewString()
		name := gofakeit.Name()
		status := gofakeit.Number(0, 1)
		payload := fmt.Sprintf(`{"name":"%s","status":%d}`, name, status)
		c, rec := m.prepareContext(strings.NewReader(payload))
		c.SetParamNames("taskId")
		c.SetParamValues(id)

		// stubs
		err := gofakeit.Error()
		updateParams := fmt.Sprintf(`{"name":"%s","status":%d,"id":"%s"}`, name, status, id)
		m.mockTaskCtl.EXPECT().Update(gomock.Any(), EqJSON(t, updateParams)).Return(nil, err)

		// assert
		err = m.handler.UpdateTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.New()
		c, rec := m.prepareContext(nil)
		c.SetParamNames("taskId")
		c.SetParamValues(id.String())

		// stubs
		m.mockTaskCtl.EXPECT().Delete(gomock.Any(), id).Return(nil)

		// assert
		err := m.handler.DeleteTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)
		require.JSONEq(t, `{"success":true}`, rec.Body.String())
	})

	t.Run("bad request", func(t *testing.T) {
		m := setup(t)

		// prepare
		c, rec := m.prepareContext(nil)
		c.SetParamNames("taskId")
		c.SetParamValues("invalid")

		// assert
		err := m.handler.DeleteTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rec.Code)
		require.Contains(t, rec.Body.String(), "invalid task id")
	})

	t.Run("not found", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.New()
		c, rec := m.prepareContext(nil)
		c.SetParamNames("taskId")
		c.SetParamValues(id.String())

		// stubs
		m.mockTaskCtl.EXPECT().Delete(gomock.Any(), id).Return(controller.ErrNotFound)

		// assert
		err := m.handler.DeleteTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("error", func(t *testing.T) {
		m := setup(t)

		// prepare
		id := uuid.New()
		c, rec := m.prepareContext(nil)
		c.SetParamNames("taskId")
		c.SetParamValues(id.String())

		// stubs
		err := gofakeit.Error()
		m.mockTaskCtl.EXPECT().Delete(gomock.Any(), id).Return(err)

		// assert
		err = m.handler.DeleteTask()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
