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

main:
	for i := min; i <= max; i++ {
		stringRep := strconv.Itoa(i)

		first, _ := strconv.Atoi(string(stringRep[0]))
		second, _ := strconv.Atoi(string(stringRep[1]))
		third, _ := strconv.Atoi(string(stringRep[2]))
		fourth, _ := strconv.Atoi(string(stringRep[3]))
		fifth, _ := strconv.Atoi(string(stringRep[4]))
		sixth, _ := strconv.Atoi(string(stringRep[5]))

		matches := map[int]int{first:1, second:1, third:1, fourth:1, fifth:1, sixth:1}

		if first > second || second > third || third > fourth || fourth > fifth || fifth > sixth {
			continue
		}

		if first == second  {
			matches[first] += 1
		}

		if second == third  {
			matches[second] += 1
		}
			
		if third == fourth  {
			matches[third] += 1
		}

		if fourth == fifth  {
			matches[fourth] += 1
		}

		if fifth == sixth  {
			matches[fifth] += 1
		}

		for _, m := range matches {
			if m == 2 {
				count +=1
				continue main
			}
		}
	}

	fmt.Println(count)
}
