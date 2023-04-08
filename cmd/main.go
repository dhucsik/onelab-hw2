package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/dhucsik/onelab-hw2/config"
	"github.com/dhucsik/onelab-hw2/service"
	"github.com/dhucsik/onelab-hw2/storage"
	"github.com/dhucsik/onelab-hw2/transport/http"
	"github.com/dhucsik/onelab-hw2/transport/http/handler"
)

func main() {
	log.Fatalln(run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}

	// storage пересекается с названием пакета
	storage, err := storage.New(ctx, conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc, svcErr := service.NewManager(storage)
	if svcErr != nil {
		return svcErr
	}

	h := handler.NewManager(conf, svc)
	HTTPServer := http.NewServer(conf, h)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
