package main

import "fmt"

func main() {
	digits1 := []int{4, 3, 9, 1}
	digits2 := []int{9, 2, 9, 9}

	out1, out2 := plusOne(digits1), plusOne(digits2)
	fmt.Println(out1, out2)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)

	return digits
}
