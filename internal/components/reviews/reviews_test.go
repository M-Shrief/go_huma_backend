package reviews

import (
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestPutReview(t *testing.T) {
	_, api := humatest.New(t)

	RegisterAPI(api)

	resp := api.Post("/reviews", map[string]any{
		"author": "daniel",
		"rating": 5,
	})

	if resp.Code != 201 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestPutReviewError(t *testing.T) {
	_, api := humatest.New(t)

	RegisterAPI(api)

	resp := api.Post("/reviews", map[string]any{
		"rating": 10,
	})

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}
