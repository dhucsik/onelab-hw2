package model

import (
	"time"
)

type User struct {
	ID      uint
	Name    string
	SurName string
	Pass    string
	Login   string
}

type CreateResp struct {
	ID        uint
	CreatedAt time.Time
}

type UpdateResp struct {
	ID        uint
	UpdatedAt time.Time
}

type DeleteResp struct {
	ID        uint
	DeletedAt time.Time
}
