package main

import (
	"go_huma_backend/router"
	"net/http"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func main() {

	// Router & API
	r := router.NewRouter()
	router.UseMiddlewares()
	router.InitAPI()

	// Start the server!
	http.ListenAndServe("127.0.0.1:3000", r)
}
