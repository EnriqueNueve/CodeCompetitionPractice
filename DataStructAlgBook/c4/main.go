/*
          __Notes__

Searching: locating information

Search problem: given an array a and integer x, find
                an integer i such that
  1. if there is no j suh that a[j] is x, then i is -1
  2. o.w., i is any j for which a[j] is x

*/


package main

import (
  "fmt"
)

func LinearSearch(key int, values []int) int {
  for i := 0; i < len(values); i++ {
    if values[i] == key {
      return i
    }
  }
  return -1
}

/*
Binary Search
Assume list is sorted least to greatest
Set mid = 0, left = 0, right = len(array) - 1
*/
func BinarySearch(key int, values []int) int {
  var mid, left, right int = 0, 0, len(values) - 1
  for left < right {
    mid = (left+right)/2
    if key > values[mid] {
      left = mid + 1
    } else {
      right = mid
    }
  }
  if values[left] == key {
    return left
  } else {
    return -1
  }
}

func main(){
  l1 := []int{1,4,17,3,90,79,4,6,81}
  fmt.Println("Values in the l1: ",l1)

  // Perform Linear Search
  var key_a int = 81
  var key_b int = 100
  fmt.Println("Where in l1 is", key_a, "?",LinearSearch(key_a,l1))
  fmt.Println("Where in l1 is", key_b, "?",LinearSearch(key_b,l1))

  // Binary Search
  fmt.Println("\n")
  l2 := []int{1,3,4,4,6,17,79,81,90}
  fmt.Println("Values in the l2: ",l2)
  key_a = 9
  key_b = 79
  fmt.Println("Where in l2 is", key_a, "?",BinarySearch(key_a,l2))
  fmt.Println("Where in l2 is", key_b, "?",BinarySearch(key_b,l2))

}
