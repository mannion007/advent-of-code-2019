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
)

type vector struct {
	x int
	y int
}

func main() {

	trace1 := trace(readFile("wire1.txt"))
	trace2 := trace(readFile("wire2.txt"))

	intersect := make(map[vector]struct{})
	for k, v := range trace1 {
		if _, ok := trace2[k]; ok {
			intersect[k] = v
		}
	}

	for k, _ := range intersect {
		k.x = int(math.Abs(float64(k.x)))
		k.y = int(math.Abs(float64(k.y)))

		if k.x+k.y < minManhattan || minManhattan == 0 {
			minManhattan = k.x + k.y
		}
	}

	fmt.Println(minManhattan)
}

func trace(dirs []string) map[vector]struct{} {
	var x, y int = 0, 0

	trace := make(map[vector]struct{})

	for _, d := range dirs {

		dist, _ := strconv.ParseInt(d[1:], 10, 64)

		for i := 0; i < int(dist); i++ {
			x += xMap[string(d[:1])]
			y += yMap[string(d[:1])]

			trace[vector{x, y}] = struct{}{}
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
