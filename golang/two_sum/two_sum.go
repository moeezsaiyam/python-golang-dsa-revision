// Two Sum – Find two numbers that add up to a target.

// Example:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

package main

import "fmt"

func two_sum(nums []int, target int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{0}
}

func two_sum_linear(nums []int, target int) []int {
	seen := make(map[int]bool)

	for _, num := range nums {
		compliment := target - num
		if seen[compliment] {
			return []int{compliment, num}
		}
		seen[num] = true
	}
	return []int{0}
}

func main() {
	nums := []int{2, 4, 5, 9}
	target := 7

	sums := two_sum(nums, target)
	fmt.Println(sums)

	sumss := two_sum_linear(nums, target)
	fmt.Println(sumss)
}
