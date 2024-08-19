package main

import (
	"context"
	"fmt"
	"go_huma_backend/internal/config"
	"go_huma_backend/internal/database"
	"go_huma_backend/router"
	"net/http"
)

func main() {
	// Load Config Variables
	config.LoadENV()

	// Database
	conn, _ := database.Connect()
	defer conn.Close(context.Background())

	// Router & API
	r := router.NewRouter()
	router.UseMiddlewares()
	router.InitAPI()

	// Start the server!
	http.ListenAndServe(
		fmt.Sprintf("%v:%v", config.HOST, config.PORT),
		r,
	)
}
