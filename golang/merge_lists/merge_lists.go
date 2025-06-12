package main

import (
	"fmt"
	"sort"
)

func merge_sorted_lists(li1 []int, li2 []int) []int {
	p1 := 0
	p2 := 0
	li := make([]int, 0, len(li1)+len(li2))

	for p1 < len(li1) && p2 < len(li2) {
		if li1[p1] < li2[p2] {
			li = append(li, li1[p1])
			p1 += 1
		} else {
			li = append(li, li2[p2])
			p2 += 1
		}
	}

	for p1 < len(li1) {
		li = append(li, li1[p1])
		p1 += 1
	}

	for p2 < len(li2) {
		li = append(li, li2[p2])
		p2 += 1
	}

	return li
}

func merge_sorted_lists_concat(li1 []int, li2 []int) []int {
	li := append(li1, li2...)
	sort.Ints(li)

	return li
}

func main() {
	li1 := []int{1, 2, 4, 6}
	li2 := []int{1, 3, 5, 6, 8}

	li := merge_sorted_lists(li1, li2)

	fmt.Println(li)

	li3 := merge_sorted_lists_concat(li1, li2)
	fmt.Println(li3)
}
