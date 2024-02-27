package main

import (
	"context"
	"github.com/ahghazey/rate_limiter/pkg/config"
	"github.com/ahghazey/rate_limiter/pkg/http/handler"
	"github.com/ahghazey/rate_limiter/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	router := handler.Handler()
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	httpServer := server.NewHttpServer(configuration, router)

	err = httpServer.Start()
	if err != nil {
		log.Fatal("Error starting server: ", err)

	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	log.Println("Got signal: ", sig)
	ctx, cancel := context.WithTimeout(ctx, configuration.WaitTimeOut())
	defer cancel()

	err = httpServer.Stop(ctx)
	if err != nil {
		log.Fatal("Error stopping server: ", err)
	}

}
