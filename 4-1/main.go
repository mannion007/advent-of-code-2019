package main

import (
	"fmt"
	"strconv"
)

var(
	numbers int	
) 

func main() {

	min, max := 234208,765869

	count := 0

	for i := min; i <= max; i++ {
		stringRep := strconv.Itoa(i)

		first, _ := strconv.Atoi(string(stringRep[0]))
		second, _ := strconv.Atoi(string(stringRep[1]))
		third, _ := strconv.Atoi(string(stringRep[2]))
		fourth, _ := strconv.Atoi(string(stringRep[3]))
		fifth, _ := strconv.Atoi(string(stringRep[4]))
		sixth, _ := strconv.Atoi(string(stringRep[5]))

		if first == second || second == third || third == fourth || fourth == fifth || fifth == sixth {
			if first >= second && second >= third && third >= fourth && fourth >= fifth && fifth >= sixth {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
