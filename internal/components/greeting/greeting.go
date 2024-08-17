package greeting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

var (
	greetings = "Greetings"
)

type GreetingInput struct {
	Name string `path:"name" maxLength:"10" example:"world" doc:"Name to greet"`
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func GreetingHandler(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}

func RegisterAPI(api huma.API) {
	// huma.Get(api, "/greeting/{name}", GreetingHandler)
	huma.Register(
		api,
		huma.Operation{
			OperationID: "get-greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
			Summary:     "Get a greeting",
			Description: "Get a greeting for a person by name.",
			Tags:        []string{greetings},
		},
		GreetingHandler,
	)
}
