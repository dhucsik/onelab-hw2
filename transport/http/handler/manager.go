package handler

import (
	"github.com/dhucsik/onelab-hw2/config"
	"github.com/dhucsik/onelab-hw2/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(conf *config.Config, srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
