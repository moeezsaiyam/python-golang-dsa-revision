package main

import "fmt"

// # The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

// # F(0) = 0, F(1) = 1
// # F(n) = F(n - 1) + F(n - 2), for n > 1.
// # Given n, calculate F(n)

func fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci_memo(n int, mem map[int]int) int {
	_, key := mem[n]
	if key {
		return mem[n]
	}

	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	mem[n] = fibonacci_memo(n-1, mem) + fibonacci_memo(n-2, mem)
	return mem[n]
}

func main() {
	fmt.Println(fibonacci(40))
	fmt.Println(fibonacci_memo(40, make(map[int]int)))
}
