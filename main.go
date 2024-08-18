package main

import (
	"go_huma_backend/router"
	"net/http"
)

func main() {

	// Router & API
	r := router.NewRouter()
	router.UseMiddlewares()
	router.InitAPI()

	// Start the server!
	http.ListenAndServe("127.0.0.1:3000", r)
}
