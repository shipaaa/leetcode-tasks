package main

import (
	"fmt"
	"math"
)

func main() {
	input1, input2, input3, input4, input5 := "42", "-042", "1337c0d3", "0-1", "words and 987"
	output1, output2, output3, output4, output5 := myAtoi(input1), myAtoi(input2), myAtoi(input3), myAtoi(input4), myAtoi(input5)
	fmt.Println(output1, output2, output3, output4, output5)

}

func myAtoi(s string) int {
	var result, startIndex int
	sign := 1

	for startIndex < len(s) && s[startIndex] == ' ' {
		startIndex++
	}

	if startIndex < len(s) {
		if s[startIndex] == '-' {
			sign = -1
			startIndex++
		} else if s[startIndex] == '+' {
			startIndex++
		}
	}

	for startIndex < len(s) {
		if s[startIndex] < '0' || s[startIndex] > '9' {
			break
		}

		digit := int(s[startIndex] - '0')

		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		startIndex++
	}

	return result * sign
}
