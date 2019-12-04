package main

import (
	"testing"
)

func TestTraceWire(t *testing.T) {

	input := []string{"R8", "U5", "L5", "D3"}
	expected := []vector{{0, 0}, {8, 0}, {8, 5}, {3, 5}, {3, 2}}

	got := traceWire(input)

	if len(got) != len(expected) {
		t.Errorf("first(%v) = %v; want %v", input, got, expected)
	}

}
