package main

import (
	"fmt"
)

func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	nums21, nums22 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}

	out1, out2 := intersect(nums1, nums2), intersect(nums21, nums22)

	fmt.Println(out1, out2)
}

func intersect(nums1 []int, nums2 []int) []int {
	var intersection []int

	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if count, exists := m[v]; exists && count > 0 {
			intersection = append(intersection, v)
			m[v]--
		}
	}

	return intersection
}
