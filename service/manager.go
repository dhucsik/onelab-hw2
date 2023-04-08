package service

import (
	"errors"

	"github.com/dhucsik/onelab-hw2/storage"
)

type Manager struct {
	User IUserService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	uSrv := NewUserService(storage)
	if storage == nil { // это лучше сделаь 1 строкой
		return nil, errors.New("no storage provided")
	}

	return &Manager{
		User: uSrv,
	}, nil
}
