package main

import (
	"fmt"
	"sort"
)

func main() {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"
	output1 := isAnagram(s1, t1)
	output2 := isAnagram(s2, t2)
	fmt.Println(output1, output2)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		runeValuesOfS []int
		runeValuesOfT []int
	)

	for i := range s {
		runeValuesOfS = append(runeValuesOfS, int(s[i]))
		runeValuesOfT = append(runeValuesOfT, int(t[i]))
	}

	sort.Ints(runeValuesOfS)
	sort.Ints(runeValuesOfT)
	for i, v := range runeValuesOfS {
		if v != runeValuesOfT[i] {
			return false
		}
	}

	return true
}
