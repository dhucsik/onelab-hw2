package postgre

import (
	"errors"
	"time"

	"github.com/dhucsik/onelab-hw2/model"
)

var id uint

type UserRepository struct {
	db []model.User
}

func (r *UserRepository) List() ([]model.User, error) {
	return r.db, nil
}

func (r *UserRepository) Create(m model.User) (model.CreateResp, error) {
	id++
	m.ID = id
	r.db = append(r.db, m)
	return model.CreateResp{
		ID:        m.ID,
		CreatedAt: time.Now(),
	}, nil
}

func (r *UserRepository) Get(ID uint) (model.User, error) {
	for _, el := range r.db {
		if el.ID == ID {
			return el, nil
		}
	}

	return model.User{}, errors.New("user not found")
}

func (r *UserRepository) Delete(ID uint) (model.DeleteResp, error) {
	for i, el := range r.db {
		if el.ID == ID {
			r.db = append(r.db[:i], r.db[i+1:]...)
		}
	}
	return model.DeleteResp{
		ID:        ID,
		DeletedAt: time.Now(),
	}, nil
}

func (r *UserRepository) Update(ID uint, m model.User) (model.UpdateResp, error) {
	for i, el := range r.db {
		if el.ID == ID {
			r.db[i] = m
			r.db[i].ID = ID
		}
	}

	return model.UpdateResp{
		ID:        ID,
		UpdatedAt: time.Now(),
	}, nil
}

func NewUserRepository() *UserRepository {
	db := make([]model.User, 0)
	return &UserRepository{db: db}
}
