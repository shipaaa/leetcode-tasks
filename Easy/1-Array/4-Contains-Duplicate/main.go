package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums3 := []int{1, 2, 3, 4}

	out1 := containsDuplicate(nums1)
	out2 := containsDuplicate(nums2)
	out3 := containsDuplicate(nums3)
	fmt.Println(out1, out2, out3)
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
		if m[v] >= 2 {
			return true
		}
	}

	return false
}
