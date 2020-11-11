package main

import (
  "fmt"
)

// Determines which of two ints are the smallest
func minInt(a int, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}

// Solves for which box would contain the most area
func maxArea(height []int) int {
  var max_area int = 0
  var temp_max_area int = 0
  var step_a int = 1
  var step_b int = 0

  for i := 0; i < len(height) - 1; i++ {
    for j := i + 1; j > 0 ; j-- {
      temp_max_area = minInt(height[step_a],height[step_b])
      temp_max_area = temp_max_area*j
      if temp_max_area > max_area {
        max_area = temp_max_area
      }
      temp_max_area = 0
      step_b += 1
    }
    step_b = 0
    step_a += 1
  }

  return max_area
}

func main() {
  // test 1
  height := []int{4,3,2,1,4}
  max_area := maxArea(height)
  fmt.Println(max_area) // Solution is 16

  // test 2
  height = []int{1,1}
  max_area = maxArea(height)
  fmt.Println(max_area) // Solution is 1

  // test 3
  height = []int{1,8,6,2,5,4,8,3,7}
  max_area = maxArea(height)
  fmt.Println(max_area) // Solution is 49
}
