package handler

import (
	"github.com/ahghazey/rate_limiter/pkg/http/rest"
	"github.com/ahghazey/rate_limiter/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// go get github.com/go-chi/chi/v5

func Handler() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recovery)
	router.Get("/health", rest.Health())
	return router
}
