package main

import "fmt"

func main() {
	strs1 := []string{"flower", "flow", "flight"}

	out1 := longestCommonPrefix(strs1)

	fmt.Println(out1)
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !startswith(strs[i], prefix) {
			fmt.Println(prefix, &prefix)
			prefix = prefix[:len(prefix)-1]
			fmt.Println(prefix, &prefix)
		}

		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

func startswith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}

	return s[:len(prefix)] == prefix
}
