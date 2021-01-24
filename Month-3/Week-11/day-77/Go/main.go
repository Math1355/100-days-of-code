package main

import (
	"fmt"
	"strings"
)

func nextMove(n, r, c int, grid [][]string) (result string) {
	var up, left int
	princess := [2]int{0, 0}

	/* find the princess */

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "p" {
				princess[0] = i
				princess[1] = j
			}
		}
	}

	up = princess[0] - r
	if up < 0 {
		result = fmt.Sprintf(strings.Repeat("DOWN ", up*-1))
	} else if up > 0 {
		result = fmt.Sprintf(strings.Repeat("UP ", up))
	}

	left = princess[1] - c
	if left < 0 {
		result = fmt.Sprintf(strings.Repeat("LEFT ", left*-1))

	} else if left > 0 {
		result = fmt.Sprintf(strings.Repeat("RIGHT ", left))
	}

	return result
}

/* Tail starts here */
func main() {

	n := 5
	r := 2
	c := 3

	grid := [][]string{{"-", "-", "-", "-", "-"}, {"-", "-", "-", "-", "-"}, {"p", "-", "-", "m", "-"}, {"-", "-", "-", "-", "-"}, {"-", "-", "-", "-", "-"}}

	fmt.Println(strings.Split(nextMove(n, r, c, grid), " "))
}
