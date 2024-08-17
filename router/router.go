package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var R *chi.Mux

func NewRouter() *chi.Mux {
	R = chi.NewRouter()
	return R
}

// Use standard middlewares.
func UseMiddlewares() {
	R.Use(middleware.Logger)
	R.Use(middleware.Recoverer)
}
