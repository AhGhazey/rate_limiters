package handler

import (
	"github.com/ahghazey/rate_limiter/pkg/http/rest"
	"github.com/ahghazey/rate_limiter/pkg/limiters"
	"github.com/ahghazey/rate_limiter/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// go get github.com/go-chi/chi/v5

func Handler(rateLimiter limiters.RateLimiter) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.TraceHeaders)
	router.Use(middleware.Recovery)
	//router.Use(middleware.TokenBucketRateLimiter(rateLimiter))
	router.Get("/health", rest.Health())
	return router
}
