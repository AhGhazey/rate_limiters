package main

import (
	"context"
	"fmt"
	clients "github.com/ahghazey/rate_limiter/pkg/clients/redis"
	"github.com/ahghazey/rate_limiter/pkg/config"
	"github.com/ahghazey/rate_limiter/pkg/http/handler"
	"github.com/ahghazey/rate_limiter/pkg/limiters"
	"github.com/ahghazey/rate_limiter/pkg/server"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configuration.RedisHost(), configuration.RedisPort()),
		Password: configuration.RedisPassword(),
		DB:       configuration.RedisDb(),
	})
	cache := clients.NewRedisClient(client)
	tokenBucketService := limiters.NewRateLimiterService(cache)
	router := handler.Handler(tokenBucketService)

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
