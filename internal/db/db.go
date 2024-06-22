package db

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrOnlyPointer   = errors.New("only pointer")
)

type Model string

const (
	Task Model = "task"
)

type Database interface {
	Get(model Model, id uuid.UUID) (interface{}, error)
	// List by create order
	List(model Model) ([]interface{}, error)
	Create(model Model, id uuid.UUID, value interface{}) error
	Update(model Model, id uuid.UUID, value interface{}) error
	Delete(model Model, id uuid.UUID) error
}

type databaseManager struct {
	database map[Model]*modelDatabase
}

type modelDatabase struct {
	dataMap map[uuid.UUID]interface{}

	orders []uuid.UUID
}

func New() Database {
	db := &databaseManager{
		database: map[Model]*modelDatabase{
			Task: {
				dataMap: map[uuid.UUID]interface{}{},
				orders:  []uuid.UUID{},
			},
		},
	}
	return db
}

func (db *databaseManager) Get(model Model, id uuid.UUID) (interface{}, error) {
	modelDB := db.database[model]
	item, ok := modelDB.dataMap[id]
	if !ok {
		return nil, ErrNotFound
	}

	return item, nil
}

func (db *databaseManager) List(model Model) ([]interface{}, error) {
	modelDB := db.database[model]

	list := make([]interface{}, 0, len(modelDB.orders))
	for _, id := range modelDB.orders {
		list = append(list, modelDB.dataMap[id])
	}

	return list, nil
}

func (db *databaseManager) Create(model Model, id uuid.UUID, value interface{}) error {
	if err := isPointer(value); err != nil {
		return err
	}

	modelDB := db.database[model]
	if _, ok := modelDB.dataMap[id]; ok {
		return ErrAlreadyExists
	}

	modelDB.dataMap[id] = value
	modelDB.orders = append(modelDB.orders, id)
	return nil
}

func (db *databaseManager) Update(model Model, id uuid.UUID, value interface{}) error {
	if err := isPointer(value); err != nil {
		return err
	}

	if _, err := db.Get(model, id); err != nil {
		return err
	}

	db.database[model].dataMap[id] = value
	return nil
}

func (db *databaseManager) Delete(model Model, id uuid.UUID) error {
	if _, err := db.Get(model, id); err != nil {
		return err
	}

	modelDB := db.database[model]
	delete(modelDB.dataMap, id)

	for i, item := range modelDB.orders {
		if item == id {
			modelDB.orders = append(modelDB.orders[:i], modelDB.orders[i+1:]...)
			break
		}
	}

	return nil
}

func isPointer(v interface{}) error {
	if v == nil {
		return ErrOnlyPointer
	}

	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return ErrOnlyPointer
	}

	return nil
}
