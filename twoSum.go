package main

import (
  "fmt"
)

func twoSum(nums []int, target int) []int {
  var results []int
  for i := 0; i < len(nums); i++ {
    for j := 0; j < len(nums); j++ {
      // Skip identical
      if i == j {
        continue
      }
      if nums[i] + nums[j] == target {
        results = append(results,[]int{i,j}...)
        return results
      }
    }
  }
  return results
}

func main() {
  x := []int{1,4,5}
  y := twoSum(x,6)
  fmt.Println(y)
}
