package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vmalits/taskboard/backend/internal/http/handler"
)

type Router struct {
	HealthHandler http.HandlerFunc
}

func New(health *handler.HealthHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/health", health)

	return r
}
