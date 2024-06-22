package store

import (
	"testing"

	mock_db "github.com/dragon-huang0403/todo-go/internal/db/mock"
	"go.uber.org/mock/gomock"
)

type testMain struct {
	store Store

	mockDB *mock_db.MockDatabase
}

func setup(t *testing.T) *testMain {
	ctl := gomock.NewController(t)
	t.Cleanup(ctl.Finish)

	mockDB := mock_db.NewMockDatabase(ctl)

	store := New(mockDB)

	return &testMain{
		store:  store,
		mockDB: mockDB,
	}
}
