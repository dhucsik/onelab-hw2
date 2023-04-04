package storage

import (
	"context"

	"github.com/dhucsik/onelab-hw2/config"
	"github.com/dhucsik/onelab-hw2/model"
	"github.com/dhucsik/onelab-hw2/storage/postgre"
)

type IUserRepository interface {
	Get(ID uint) (model.User, error)
	Create(req model.User) (model.CreateResp, error)
	List() ([]model.User, error)
	Update(ID uint, m model.User) (model.UpdateResp, error)
	Delete(ID uint) (model.DeleteResp, error)
}

type Storage struct {
	User IUserRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	uRepo := postgre.NewUserRepository()

	var storage Storage
	storage.User = uRepo
	return &storage, nil
}
