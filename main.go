package main

import (
	"go_huma_backend/router"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func main() {
	// Router & API
	r := router.NewRouter()
	router.UseMiddlewares()
	router.InitAPI(r, huma.DefaultConfig("My API", "1.0.0"))
	router.RegisterAPIs()

	// Start the server!
	http.ListenAndServe("127.0.0.1:3000", r)
}
