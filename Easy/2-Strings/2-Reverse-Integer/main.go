package main

import (
	"fmt"
	"math"
)

func main() {
	input1, input2, input3 := 123, -123, 1534236469
	output1, output2, output3 := reverse(input1), reverse(input2), reverse(input3)
	fmt.Println(output1, output2, output3) // 321 -321 0
	fmt.Println(123 / 10)
}

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result > math.MaxInt32 || result < -math.MaxInt32 {
		return 0
	}

	return result
}
