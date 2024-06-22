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
	List(model Model) ([]interface{}, error)
	Create(model Model, id uuid.UUID, value interface{}) error
	Update(model Model, id uuid.UUID, value interface{}) error
	Delete(model Model, id uuid.UUID) error
}

type db struct {
	database map[Model]map[uuid.UUID]interface{}
}

func New() Database {
	return &db{
		database: map[Model]map[uuid.UUID]interface{}{},
	}
}

func (db *db) Get(model Model, id uuid.UUID) (interface{}, error) {
	if _, ok := db.database[model]; !ok {
		return nil, ErrNotFound
	}

	if _, ok := db.database[model][id]; !ok {
		return nil, ErrNotFound
	}

	return db.database[model][id], nil
}

// TODO: Implement sorting
func (db *db) List(model Model) ([]interface{}, error) {
	if _, ok := db.database[model]; !ok {
		return nil, ErrNotFound
	}

	var list []interface{}
	for _, v := range db.database[model] {
		list = append(list, v)
	}

	return list, nil
}

func (db *db) Create(model Model, id uuid.UUID, value interface{}) error {
	if err := isPointer(value); err != nil {
		return err
	}

	if _, ok := db.database[model]; !ok {
		db.database[model] = map[uuid.UUID]interface{}{}
	}

	if _, ok := db.database[model][id]; ok {
		return ErrAlreadyExists
	}

	db.database[model][id] = value
	return nil
}

func (db *db) Update(model Model, id uuid.UUID, value interface{}) error {
	if err := isPointer(value); err != nil {
		return err
	}

	if _, err := db.Get(model, id); err != nil {
		return err
	}

	db.database[model][id] = value
	return nil
}

func (db *db) Delete(model Model, id uuid.UUID) error {
	if _, err := db.Get(model, id); err != nil {
		return err
	}

	delete(db.database[model], id)
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
