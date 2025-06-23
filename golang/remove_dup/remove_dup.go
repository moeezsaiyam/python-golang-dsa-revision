package main

import "fmt"

func remove_duplicate[T comparable](li []T) []T {
	checked := make(map[T]bool)
	opt := []T{}

	for _, val := range li {
		if !checked[val] {
			checked[val] = true
			opt = append(opt, val)
		}
	}
	return opt

}

func main() {
	li := []int{2, 3, 2, 2, 4, 5, 6, 7}
	fmt.Println(remove_duplicate(li))
}
