package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	nums2 := []int{0}

	moveZeroes(nums1)
	moveZeroes(nums2)

	fmt.Println(nums1, nums2)
}

func moveZeroes(nums []int) {
	var lastNonZero int

	for i, num := range nums {
		if num != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
