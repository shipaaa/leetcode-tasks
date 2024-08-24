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
	var interimS []byte
	for i := len(s) - 1; i >= 0; i-- {
		interimS = append(interimS, s[i])
	}
	for i := range s {
		s[i] = interimS[i]
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
