package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	// Declare lenof passed string length
	var n int = len(s)

	// Split string into slice of strings for ease of looping
	s_split := strings.Split(s, "")

	// Declare map to stroe slice of each cols letters
	m := make(map[int][]string)

	// Logic to traverse zigzag of passed string
	i := 0
	col := 0
	for i < len(s_split) {
		if col%2 == 0 {
			// Case: full column
			for j := 0; j < numRows && i < n; j++ {
				m[j] = append(m[j], s_split[i])
				i++
			}
		} else {
			// Case: single value colummn of zigzag
			for j := numRows - 2; j > 0 && i < n; j-- {
				m[j] = append(m[j], s_split[i])
				i++
			}
		}
		col++
	}

	// Combine individual slices in map to make single string
	var new_string string
	var total_slice []string
	for i := 0; i < numRows; i++ {
		total_slice = append(total_slice, m[i]...)
	}
	new_string = strings.Join(total_slice[:], "")

	return new_string
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	sc := convert(s, numRows)
	fmt.Println(sc)

	s = "PAYPALISHIRING"
	numRows = 4
	sc = convert(s, numRows)
	fmt.Println(sc)

	s = "A"
	numRows = 1
	sc = convert(s, numRows)
	fmt.Println(sc)
}
