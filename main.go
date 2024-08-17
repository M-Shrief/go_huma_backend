package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a new router & API
	router := chi.NewMux()
	router.Use(middleware.Logger)
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// // Register GET /greeting/{name} handler.
	// huma.Get(
	// 	api,
	// 	"/greeting/{name}",
	// 	func(
	// 		ctx context.Context,
	// 		input *struct {
	// 			Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	// 		},
	// 	) (*GreetingOutput, error) {
	// 		resp := &GreetingOutput{}
	// 		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	// 		return resp, nil
	// 	},
	// )

	// Register GET /greeting/{name} handler with more information for the docs.
	huma.Register(
		api,
		// Use huma.Operation to add more information for the docs.
		huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "Get a greeting",
			Description: "Get a greeting for a person by name.",
			Tags:        []string{"Greetings"},
		},
		func(
			ctx context.Context,
			input *struct {
				Name string `path:"name" maxLength:"10" example:"world" doc:"Name to greet"`
			},
		) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		},
	)

	// Start the server!
	http.ListenAndServe("127.0.0.1:3000", router)
}
