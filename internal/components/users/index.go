package users

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID:   "signup-users",
			Method:        http.MethodPost,
			Path:          "/api/users/signup",
			Summary:       "User Signup",
			Description:   "User signing up.",
			Tags:          []string{"Users"},
			DefaultStatus: http.StatusCreated,
		},
		SignupHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "login-users",
			Method:        http.MethodPost,
			Path:          "/api/users/login",
			Summary:       "User login",
			Description:   "User logining in.",
			Tags:          []string{"Users"},
			DefaultStatus: http.StatusAccepted,
		},
		LoginHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "update-users",
			Method:        http.MethodPut,
			Path:          "/api/users/me",
			Summary:       "Update User",
			Description:   "Update user's data.",
			Tags:          []string{"Users"},
			DefaultStatus: http.StatusAccepted,
			Security: []map[string][]string{
				{"bearer": {}},
			},
		},
		UpdateHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "delete-users",
			Method:        http.MethodDelete,
			Path:          "/api/users/me",
			Summary:       "Delete User",
			Description:   "Delete user account.",
			Tags:          []string{"Users"},
			DefaultStatus: http.StatusAccepted,
			Security: []map[string][]string{
				{"bearer": {}},
			},
		},
		DeleteHandler,
	)
}
