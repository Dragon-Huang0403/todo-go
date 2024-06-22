package controller

import (
	"testing"

	mock_store "github.com/dragon-huang0403/todo-go/internal/store/mock"
	"go.uber.org/mock/gomock"
)

type testMain struct {
	controller *Controller

	mockStore *mock_store.MockStore
}

func setup(t *testing.T) *testMain {
	ctl := gomock.NewController(t)
	t.Cleanup(ctl.Finish)

	mockStore := mock_store.NewMockStore(ctl)

	controller := New(mockStore)

	return &testMain{
		controller: controller,
		mockStore:  mockStore,
	}
}
