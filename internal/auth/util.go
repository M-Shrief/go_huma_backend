package auth

import (
	"golang.org/x/exp/slices"
)

func isAuthorized(onlyAuthorizedFor, permissions []string) bool {
	var isAuthorized bool
	for _, permission := range permissions {
		isAuthorized = slices.Contains(onlyAuthorizedFor, permission) || false
	}
	return isAuthorized
}
