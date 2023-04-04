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
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	return &Manager{
		User: uSrv,
	}, nil
}
