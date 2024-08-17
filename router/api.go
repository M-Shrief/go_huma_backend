package router

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

var API huma.API

func InitAPI(router *chi.Mux, config huma.Config) huma.API {
	API = humachi.New(router, config)
	return API
}
