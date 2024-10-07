package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 1, 1}

	out1, out2 := threeSum(nums1), threeSum(nums2)

	fmt.Println(out1, out2)
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left += 1
				right -= 1

				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			} else if total < 0 {
				left += 1
			} else if total > 0 {
				right -= 1
			}
		}
	}

	return result
}
