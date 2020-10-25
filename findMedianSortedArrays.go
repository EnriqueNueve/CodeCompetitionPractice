package main

import (
  "fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  m := len(nums1)
  n := len(nums2)
  o := m+n

  // Special case of o == 1
  if o == 1 {
    if m == 0 {
      return float64(nums2[0])
    } else {
      return float64(nums1[0])
    }
  }

  mid := (o/2)+1

  var vals []int
  for i := 0; i < mid; i++ {
    if len(nums1) != 0 && len(nums2) != 0 {
      if nums1[0] > nums2[0] {
        vals = append(vals,nums2[0])
        if len(nums2) == 1 {
          nums2 = []int{}
        } else {
          nums2 = nums2[1:]
        }
      } else {
        vals = append(vals,nums1[0])
        if len(nums1) == 1 {
          nums1 = []int{}
        } else {
          nums1 = nums1[1:]
        }
      }
    } else if len(nums1) > 0 && len(nums2) == 0 {
      vals = append(vals,nums1[0])
      nums1 = nums1[1:]
    } else if len(nums2) > 0 && len(nums1) == 0 {
      vals = append(vals,nums2[0])
      nums2 = nums2[1:]
    }
  }


  if o%2 == 0{
    return (float64(vals[len(vals)-1])+float64(vals[len(vals)-2]))/2
  } else {
    return float64(vals[len(vals)-1])
  }
}


func main() {
  nums1 := []int{}
  nums2 := []int{1,2}
  median := findMedianSortedArrays(nums1,nums2)
  fmt.Println(median) // 1.5

  nums1 = []int{1,4}
  nums2 = []int{2,3}
  median = findMedianSortedArrays(nums1,nums2)
  fmt.Println(median) // 2.5

  nums1 = []int{0,0}
  nums2 = []int{0,0}
  median = findMedianSortedArrays(nums1,nums2)
  fmt.Println(median) // 0

  nums1 = []int{2}
  nums2 = []int{}
  median = findMedianSortedArrays(nums1,nums2)
  fmt.Println(median) // 2

  nums1 = []int{1,3}
  nums2 = []int{2}
  median = findMedianSortedArrays(nums1,nums2)
  fmt.Println(median) // 2
}
