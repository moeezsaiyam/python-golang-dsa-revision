package main

import "fmt"

func is_palindrome(p string) bool {
	r := []rune(p)

	for i := range len(p) / 2 {
		if r[i] != r[len(p)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(is_palindrome("abac"))
}
