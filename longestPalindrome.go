package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := fromMid(s, i, i)
		len2 := fromMid(s, i, i+1)
		l := max(len1, len2)
		if l > end-start {
			start = i - (l-1)/2 // when l == 2, i remain unchanged, end plus + 1
			end = i + l/2       // when l is even num, i is not mid index, so i can not direct minus l/2, l has to minus one first.
		}
	}
	return s[start : end+1]
}

func fromMid(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1 // after left-- && right++, condition not met
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	w1 := "aba"
	s1 := longestPalindrome(w1)
	fmt.Println(s1)
}
