package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-open-api/internal/settings"
	"test-open-api/rest/handlers"
	"test-open-api/rest/routers"
)

//go:generate oapi-codegen -o ../internal/api/types.gen.go -package api -generate types ../swagger.yaml
//go:generate oapi-codegen -o ../internal/api/server.gen.go -package api -generate chi-server ../swagger.yaml
//go:generate oapi-codegen -o ../internal/api/spec.gen.go -package api -generate spec ../swagger.yaml

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	mainCtx := context.Background()

	cfg := settings.GetConfig()

	server := handlers.NewServer()
	router := routers.New(server)

	httpServer := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	serverStopSignal := make(chan os.Signal, 1)
	signal.Notify(serverStopSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			serverStopSignal <- syscall.SIGTERM
		}
	}()

	<-serverStopSignal

	if err := httpServer.Shutdown(mainCtx); err != nil {
		fmt.Println("Server finished")
	}

	return nil
}
