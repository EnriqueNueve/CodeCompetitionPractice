package main

import (
  "fmt"
  "strings"
)

func contains(s []string, e string) bool {
   for _, a := range s {
      if a == e {
         return true
      }
   }
   return false
}

func lengthOfLongestSubstring(s string) int {
  var max_len int = 0
  letters := strings.Split(s, "")

  // Special case of length 1
  if len(s) == 1 {
    return 1
  }

  sub_string := []string{}
  for i := 1; i < len(s); i++ {
    sub_string = append(sub_string,letters[i-1])
    for j := i; j < len(s) ; j++ {
      if contains(sub_string, letters[j]) {
        if max_len < len(sub_string) {
          max_len = len(sub_string)
        }
        sub_string = []string{}
        break
      } else {
        sub_string = append(sub_string,letters[j])
      }
    }
    if max_len < len(sub_string) {
      max_len = len(sub_string)
    }
    sub_string = []string{}
  }
  return max_len
}

func main() {
  var s = "abca defg"
  val := lengthOfLongestSubstring(s)
  fmt.Println(val)
}
