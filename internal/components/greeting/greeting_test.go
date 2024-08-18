package greeting

import (
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestGreeting(t *testing.T) {
	_, api := humatest.New(t)

	RegisterAPI(api)

	resp := api.Get("/greeting/world")
	if !strings.Contains(resp.Body.String(), "Hello, world!") {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}
