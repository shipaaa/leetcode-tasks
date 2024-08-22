package main

import (
	"fmt"
)

func main() {
	input1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	input2 := []int{1, 1, 2}

	output1 := removeDuplicates(input1) // 5, nums = [0,1,2,3,4,_,_,_,_,_]
	output2 := removeDuplicates(input2) // 2, nums = [1,2,_]

	fmt.Printf("%d, %v\n", output1, input1)
	fmt.Printf("%d, %v", output2, input2)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
