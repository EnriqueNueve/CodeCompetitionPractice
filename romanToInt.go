package main

import (
	"fmt"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func romanToInt(s string) int {
	romanNumeralDict := map[string]int{
		"M":  1000,
		"D":  500,
		"C":  100,
		"L":  50,
		"X":  10,
		"V":  5,
		"I":  1,
	}

	var vals []int
	var head int = 0
	s_split := strings.Split(s, "")
	for head < len(s_split) {
		if s_split[head] == "I" {
			// Case 1: letter in string is I
			if head+1 < len(s_split) {
				if s_split[head+1] == "V" {
					vals = append(vals, 4)
					head += 2
				} else if s_split[head+1] == "X" {
					vals = append(vals, 9)
					head += 2
				} else {
					vals = append(vals, 1)
					head += 1
				}
			} else {
				vals = append(vals, 1)
				head += 1
			}
		} else if s_split[head] == "X" {
			// Case 2: letter in string is X
			if head+1 < len(s_split) {
				if s_split[head+1] == "L" {
					vals = append(vals, 40)
					head += 2
				} else if s_split[head+1] == "C" {
					vals = append(vals, 90)
					head += 2
				} else {
					vals = append(vals, 10)
					head += 1
				}
			} else {
				vals = append(vals, 10)
				head += 1
			}
		} else if s_split[head] == "C" {
			// Case 3: letter in string is C
			if head+1 < len(s_split) {
				if s_split[head+1] == "D" {
					vals = append(vals, 400)
					head += 2
				} else if s_split[head+1] == "M" {
					vals = append(vals, 900)
					head += 2
				} else {
					vals = append(vals, 100)
					head += 1
				}
			} else {
				vals = append(vals, 100)
				head += 1
			}
		} else {
			vals = append(vals, romanNumeralDict[s_split[head]])
			head += 1
		}
	}

	n := sum(vals)

	return n
}

func main() {
	s := "III"
	val := romanToInt(s)
	fmt.Println(val)

	s = "IV"
	val = romanToInt(s)
	fmt.Println(val)

	s = "IX"
	val = romanToInt(s)
	fmt.Println(val)

	s = "LVIII"
	val = romanToInt(s)
	fmt.Println(val)

	s = "MCMXCIV"
	val = romanToInt(s)
	fmt.Println(val)
}
