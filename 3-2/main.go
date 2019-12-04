package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	xMap         = map[string]int{"U": 0, "D": 0, "L": -1, "R": 1}
	yMap         = map[string]int{"U": 1, "D": -1, "L": 0, "R": 0}
	minManhattan int
	minSteps 	 int
)

type vector struct {
	x int
	y int
}

func main() {

	trace1 := trace(readFile("wire1.txt"))
	trace2 := trace(readFile("wire2.txt"))

	intersect := make(map[vector]int)
	for k, v := range trace1 {
		if v2, ok := trace2[k]; ok {
			intersect[k] = v + v2
		}
	}

	for k, v := range intersect {
		k.x = int(math.Abs(float64(k.x)))
		k.y = int(math.Abs(float64(k.y)))

		if minManhattan == 0 || k.x+k.y < minManhattan {
			minManhattan = k.x + k.y
		}

		if minSteps == 0 || v < minSteps {
			minSteps = v
		}
	}

	fmt.Printf("Min manhattan: %v\n", minManhattan)
	fmt.Printf("Min steps: %v\n", minSteps)
}

func trace(dirs []string) map[vector]int {
	var x, y int = 0, 0

	trace := make(map[vector]int)

	steps := 1

	for _, d := range dirs {

		dist, _ := strconv.ParseInt(d[1:], 10, 64)

		for i := 0; i < int(dist); i++ {
			x += xMap[string(d[:1])]
			y += yMap[string(d[:1])]

			trace[vector{x, y}] = steps

			steps += 1
		}
	}

	return trace
}

func readFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return strings.SplitN(string(data), ",", -1)
}
