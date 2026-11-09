package json_test

import (
	"context"
	"errors"
	"testing"

	"zeroparse.dev/alpha-parsers/json"
)

// TestBudgetControls checks that the parser, generated with limit controls, stops at a Compute limit
// and at a cancelled context, and that the same input parses without them.
func TestBudgetControls(t *testing.T) {
	const input = "[1, 2, 3]"
	if _, err := json.Parse(input); err != nil {
		t.Fatal(err)
	}
	_, err := json.Parse(input, json.WithLimits(json.Limits{Compute: 1}))
	var limit *json.LimitError
	if !errors.Is(err, json.ErrLimit) || !errors.As(err, &limit) || limit.Kind != json.LimitCompute {
		t.Fatalf("budget limit: got %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := json.Parse(input, json.WithContext(ctx)); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled context: got %v", err)
	}
}
