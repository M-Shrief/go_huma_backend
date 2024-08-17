package main

import (
	"context"
	"fmt"
	"go_huma_backend/router"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	r := router.NewRouter()
	router.UseMiddlewares()
	api := router.InitAPI(r, huma.DefaultConfig("My API", "1.0.0"))

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
	http.ListenAndServe("127.0.0.1:3000", r)
}
