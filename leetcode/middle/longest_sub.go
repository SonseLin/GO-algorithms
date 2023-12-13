package main

import (
	"fmt"
)

type str string

func (s str) longestSubStr() int {
	lastOccurred := make(map[rune]int)
	current := 0
	maxLength := 0

	for i := 0; i < len(s); i++ {
		if symbol, ok := lastOccurred[rune(s[i])]; ok && symbol >= current {
			current = symbol + 1
		}

		lastOccurred[rune(s[i])] = i

		if i-current+1 > maxLength {
			maxLength = i - current + 1
		}
	}
	return maxLength
}

func main() {
	testCases := []str{
		"abcabcbb",
		"bbbbb",
		"pwwwkew",
		"aab",
		"dfdv",
	}
	for tc := range testCases {
		fmt.Println(testCases[tc].longestSubStr())
	}
}
