package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func factorial_simple(n int) int {
	fact := 1
	for i := range n {
		fact *= i + 1
	}
	return fact
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial_simple(4))
}
