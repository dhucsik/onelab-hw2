package service

import (
	"github.com/dhucsik/onelab-hw2/model"
	"github.com/dhucsik/onelab-hw2/storage"
)

type IUserService interface {
	Get(ID uint) (model.User, error)
	Create(req model.User) (model.CreateResp, error)
	List() ([]model.User, error)
	Update(ID uint, m model.User) (model.UpdateResp, error)
	Delete(ID uint) (model.DeleteResp, error)
}

type UserService struct {
	Repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Get(ID uint) (model.User, error) {
	return s.Repo.User.Get(ID)
}

func (s *UserService) List() ([]model.User, error) {
	return s.Repo.User.List()
}

func (s *UserService) Update(ID uint, m model.User) (model.UpdateResp, error) {
	return s.Repo.User.Update(ID, m)
}

func (s *UserService) Create(req model.User) (model.CreateResp, error) {
	return s.Repo.User.Create(req)
}

func (s *UserService) Delete(ID uint) (model.DeleteResp, error) {
	return s.Repo.User.Delete(ID)
}
