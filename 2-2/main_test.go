package main

import (
	"testing"
)

func TestCompute(t *testing.T) {

	// 1,9,10,3,2,3,11,0,99,30,40,50 becomes 3500,9,10,70,2,3,11,0,99,30,40,50
	// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2)
	// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6)
	// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801)
	// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99

	var examples = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, example := range examples {

		compute(&example.input)

		for k, v := range example.input {
			if example.expected[k] != v {
				t.Errorf("Error ExecIntCode expected %v; Got %v", example.expected, example.input)
				break
			}
		}
	}
}
