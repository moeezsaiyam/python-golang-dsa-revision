package main

import (
	"fmt"
	"slices"
)

func valid_anagram(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)

	slices.Sort(r1)
	slices.Sort(r2)

	s1 = string(r1)
	s2 = string(r2)

	if s1 == s2 {
		return true
	}

	return false
}

func main() {
	fmt.Println(valid_anagram("okko", "kkoo"))
}
