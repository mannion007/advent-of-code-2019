package main

import (
	"testing"
)

func TestCalcFuel(t *testing.T) {

	//For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
	//For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
	//For a mass of 1969, the fuel required is 654.
	//For a mass of 100756, the fuel required is 33583.

	examples := make(map[int]int)
	examples[12] = 2
	examples[14] = 2
	examples[1969] = 654
	examples[100756] = 33583

	for input, expected := range examples {
		got := calcFuel(input)

		if got != expected {
			t.Errorf("calcFuel(%d) = %d; want %d", input, got, expected)
		}
	}
}
