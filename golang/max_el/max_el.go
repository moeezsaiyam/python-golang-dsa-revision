package main

import (
	"fmt"
	"math"
)

func max_el_in_list(li []int) int {
	maxx := math.MinInt32

	for _, el := range li {
		if maxx < el {
			maxx = el
		}
	}
	return maxx
}

func main() {
	fmt.Println(max_el_in_list([]int{2, 3, -6, -10, 20, 5}))
}
