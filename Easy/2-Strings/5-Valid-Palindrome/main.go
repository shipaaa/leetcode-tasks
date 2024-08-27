package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input1 := "A man, a plan, a canal: Panama"
	input2 := "race a car"
	input3 := " "
	input4 := "Аргентина манит негра"
	output1, output2, output3, output4 := isPalindrome(input1), isPalindrome(input2), isPalindrome(input3), isPalindrome(input4)
	fmt.Println(output1, output2, output3, output4)
}

func isPalindrome(s string) bool {
	s = regexp.MustCompile(`[^\p{L}\p{N}]+`).ReplaceAllString(strings.ToLower(s), "")

	c := len(s) - 1
	for i := 0; i < len(s); i++ {
		fmt.Println(s[c], s[i])
		if s[c] != s[i] {
			return false
		}
		c--
	}

	return true
}
