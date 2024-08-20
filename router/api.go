package router

import (
	"go_huma_backend/internal/components/users"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

var API huma.API

func InitAPI() huma.API {
	API = humachi.New(R, huma.DefaultConfig("My API", "0.0.1"))

	registerAPIs()

	return API
}

func registerAPIs() {
	users.RegisterAPI(API)
}
