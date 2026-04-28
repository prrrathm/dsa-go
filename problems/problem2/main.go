// longest substring without a duplicate

package main

import (
	"fmt"
	"strings"
)

func containsDuplicates(s string) bool {
	if len(s) <= 1 {
		return false
	}

	seen := make(map[byte]bool)
	for i := range len(s) {
		if seen[s[i]] {
			return true
		}
		seen[s[i]] = true
	}
	return false
}
func isInString(s string, letter string) bool {
	if len(s) < 1 {
		return false
	}
	flag := strings.Contains(s, letter)
	return flag
}

func resetSubstring(s string, letter string) string {
	index := strings.Index(s, letter)
	substr := s[index+1 : len(s)-1]
	return substr
}

func lengthOfLongestSubstring(s string) int {
	if len(s) > 0 {
		substr := ""
		maxlen := 1
		for i := range len(s) {
			fmt.Println("sub string", substr, string(s[i]))
			if isInString(substr, string(s[i])) {
				fmt.Println("contains duplicates", substr)
				substr = resetSubstring(substr, string(s[i]))
			}
			substr += string(s[i])
			if len(substr) > maxlen {
				maxlen = len(substr)
			}
		}
		return maxlen
	}
	return 0
}

func main() {
	s := "dvdf"

	fmt.Println(lengthOfLongestSubstring(s))
	// fmt.Println(containsDuplicates(s))
}
