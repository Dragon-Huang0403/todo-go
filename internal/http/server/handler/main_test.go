package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dragon-huang0403/todo-go/internal/controller"
	mock_controller "github.com/dragon-huang0403/todo-go/internal/controller/mock"
	"github.com/dragon-huang0403/todo-go/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testMain struct {
	handler *Handler

	mockTaskCtl *mock_controller.MockTask
}

func setup(t *testing.T) *testMain {
	ctl := gomock.NewController(t)
	t.Cleanup(ctl.Finish)

	mockTaskCtl := mock_controller.NewMockTask(ctl)

	controller := &controller.Controller{
		Task: mockTaskCtl,
	}

	return &testMain{
		handler:     New(controller),
		mockTaskCtl: mockTaskCtl,
	}
}

func (m *testMain) prepareContext(requestBody io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validator.New()

	req := httptest.NewRequest(http.MethodPost, "/", requestBody)
	if requestBody != nil {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

type eqJSON struct {
	t     *testing.T
	value string
}

func (e *eqJSON) Matches(x interface{}) bool {
	b1, err := json.Marshal(x)
	if err != nil {
		return false
	}

	// Ignore case for simple comparison
	got := strings.ToLower(string(b1))
	want := strings.ToLower(e.value)
	return assert.JSONEq(e.t, want, got)
}

func (e *eqJSON) String() string {
	return fmt.Sprintf("eqJSON: is equal to %v", e.value)
}

func EqJSON(t *testing.T, v string) gomock.Matcher {
	return &eqJSON{
		t:     t,
		value: v,
	}
}
