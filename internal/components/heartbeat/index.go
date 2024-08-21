package heartbeat

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type HeartbeatOutput struct {
	Status int
}

type PingOutput struct {
	Body   string
	Status int
}

func RegisterAPI(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID:   "index-hearbeat",
			Method:        http.MethodGet,
			Path:          "/",
			Summary:       "Server health check",
			Description:   "Checking the server's health.",
			Tags:          []string{"Heartbeat"},
			DefaultStatus: http.StatusOK,
		},
		func(ctx context.Context, input *struct{}) (*HeartbeatOutput, error) {
			return &HeartbeatOutput{http.StatusOK}, nil
		},
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "ping-hearbeat",
			Method:        http.MethodGet,
			Path:          "/ping",
			Summary:       "Ping the server",
			Description:   "Ping the server to check it's health.",
			Tags:          []string{"Heartbeat"},
			DefaultStatus: http.StatusOK,
		},
		func(ctx context.Context, input *struct{}) (*PingOutput, error) {
			return &PingOutput{"Pong", http.StatusOK}, nil
		},
	)
}
