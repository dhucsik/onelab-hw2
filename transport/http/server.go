package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dhucsik/onelab-hw2/config"
	"github.com/dhucsik/onelab-hw2/transport/http/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(fmt.Sprintf(":%d", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
