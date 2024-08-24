package main

import (
	"fmt"
)

func main() {
	input1 := formatInput([]string{"h", "e", "l", "l", "o"})
	input2 := formatInput([]string{"H", "a", "n", "n", "a", "h"})

	fmt.Println(input1, input2)
	reverseString(input1)
	reverseString(input2)
	fmt.Println(input1, input2)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		fmt.Println(left, right)
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func formatInput(input1 []string) []byte {
	formattedInput := make([]byte, len(input1))
	for i, s := range input1 {
		if len(s) > 0 {
			formattedInput[i] = s[0]
		}
	}

	return formattedInput
}
