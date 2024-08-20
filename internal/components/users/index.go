package users

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID: "signup-users",
			Method:      http.MethodPost,
			Path:        "/users/signup",
			Summary:     "User Signup",
			Description: "User signing up.",
			Tags:        []string{"Users"},
		},
		SignupHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID: "login-users",
			Method:      http.MethodPost,
			Path:        "/users/login",
			Summary:     "User login",
			Description: "User logining in.",
			Tags:        []string{"Users"},
		},
		LoginHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID: "update-users",
			Method:      http.MethodPut,
			Path:        "/users/me",
			Summary:     "Update User",
			Description: "Update user's data.",
			Tags:        []string{"Users"},
			Security: []map[string][]string{
				{"JWT-Auth": {"JWT authentication"}},
			},
		},
		UpdateHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID: "delete-users",
			Method:      http.MethodDelete,
			Path:        "/users/me",
			Summary:     "Delete User",
			Description: "Delete user account.",
			Tags:        []string{"Users"},
			Security: []map[string][]string{
				{"JWT-Auth": {}},
			},
		},
		DeleteHandler,
	)
}
