package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()

		// assert
		v, err := db.Get(Task, id)
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrNotFound)
	})
}

func TestList(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db := New()

		// prepare
		n := gofakeit.Number(1, 2)
		values := make([]interface{}, 0, n)
		for range n {
			id := uuid.New()
			v := gofakeit.Map()
			err := db.Create(Task, id, &v)
			require.NoError(t, err)
			values = append(values, &v)
		}

		// assert
		values2, err := db.List(Task)
		require.NoError(t, err)
		require.Len(t, values2, n)
		for i, v := range values2 {
			require.Equal(t, values[i], v)
		}
	})

	t.Run("no rows", func(t *testing.T) {
		db := New()

		// assert
		values, err := db.List(Task)
		require.NoError(t, err)
		require.NotNil(t, values)
		require.Len(t, values, 0)
	})
}

func TestCreate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()

		// assert
		err := db.Create(Task, id, &value)
		require.NoError(t, err)

		v, err := db.Get(Task, id)
		require.NoError(t, err)
		require.Equal(t, &value, v)
	})

	t.Run("only pointer", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()

		// assert
		err := db.Create(Task, id, value)
		require.ErrorIs(t, err, ErrOnlyPointer)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()
		err := db.Create(Task, id, &value)
		require.NoError(t, err)

		// assert
		newValue := gofakeit.Map()
		err = db.Update(Task, id, &newValue)
		require.NoError(t, err)

		v, err := db.Get(Task, id)
		require.NoError(t, err)
		require.Equal(t, &newValue, v)
	})

	t.Run("not found", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()

		// assert
		err := db.Update(Task, id, &value)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("only pointer", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()

		// assert
		err := db.Update(Task, id, value)
		require.ErrorIs(t, err, ErrOnlyPointer)
	})
}

func TestDelete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()
		value := gofakeit.Map()
		err := db.Create(Task, id, &value)
		require.NoError(t, err)

		// assert
		err = db.Delete(Task, id)
		require.NoError(t, err)

		v, err := db.Get(Task, id)
		require.ErrorIs(t, err, ErrNotFound)
		require.Nil(t, v)
	})

	t.Run("not found", func(t *testing.T) {
		db := New()

		// prepare
		id := uuid.New()

		// assert
		err := db.Delete(Task, id)
		require.ErrorIs(t, err, ErrNotFound)
	})
}
