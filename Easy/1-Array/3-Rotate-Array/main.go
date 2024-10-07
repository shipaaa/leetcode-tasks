package main

import "fmt"

func main() {
	nums1, k1 := []int{1, 2, 3, 4, 5, 6, 7}, 12
	nums2, k2 := []int{-1, -100, 3, 99}, 2

	rotate(nums1, k1)
	rotate(nums2, k2)

	fmt.Println(nums1, nums2)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(0, n-1, nums) // Разворот всего массива
	reverse(0, k-1, nums) // Разворот первых k элементов
	reverse(k, n-1, nums) // Разворот оставшихся элементов
}

func reverse(left int, right int, nums []int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
