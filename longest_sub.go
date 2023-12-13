package main

import "fmt"

type str string

func (s str) longestSubStr() int {
	max_v, curr := 0, 0
	var sub []rune
	for _, v := range s {
		if curr == 0 {
			sub = []rune{}
			sub = append(sub, v)
			curr++
		} else {
			curr++
		}
		if curr > max_v {
			max_v = curr
		}
	}
	return max_v
}

func main() {
	test_cases := []str{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	for tc := range test_cases {
		fmt.Println(test_cases[tc].longestSubStr())
	}
}
