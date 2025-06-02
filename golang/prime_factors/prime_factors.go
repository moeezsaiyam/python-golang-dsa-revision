package main

import "fmt"

func prime_factors(num int) []int {

	factors := []int{}

	for num > 1 {
		if num%2 == 0 {
			factors = append(factors, 2)
			num = num / 2
		} else if num%3 == 0 {
			factors = append(factors, 3)
			num = num / 3
		} else {
			factors = append(factors, num)
			break
		}
	}
	return factors
}

func main() {
	factors := prime_factors(5)
	fmt.Println(factors)
}
