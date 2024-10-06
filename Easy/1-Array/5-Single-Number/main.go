package main

import "fmt"

func main() {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}

	out1, out2, out3 := singleNumber(nums1), singleNumber(nums2), singleNumber(nums3)
	fmt.Println(out1, out2, out3)
}

func singleNumber(nums []int) int {
	var result int
	interimM := make(map[int]int, len(nums))

	for _, num := range nums {
		interimM[num]++
	}

	for num, count := range interimM {
		if count == 1 {
			result = num
			break
		}
	}

	return result
}
