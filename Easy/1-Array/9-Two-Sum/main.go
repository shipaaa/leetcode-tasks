package main

import (
	"fmt"
)

func main() {
	nums, target := []int{3, 2, 4}, 6

	out := twoSum(nums, target)

	fmt.Println(out)
}

func twoSum(nums []int, target int) []int {
	interimM := make(map[int]int, len(nums))

	for i, num := range nums {
		requiredValue := target - num
		if idx, found := interimM[requiredValue]; found {
			return []int{idx, i}
		}
		interimM[num] = i
	}

	return nil
}
