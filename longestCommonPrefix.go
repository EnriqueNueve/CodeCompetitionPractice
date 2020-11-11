package main

import (
  "fmt"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := []byte{}
    baseLength, strLength := len(strs[0]), len(strs)
    for i := 0; i < baseLength; i++ {
        for j := 1; j < strLength; j++ {
            if len(strs[j]) <= i || strs[j][i] != strs[0][i]  {
                return string(prefix)
            }
        }
        prefix = append(prefix, strs[0][i])
    }
    return string(prefix)
}

func main () {
  strs := []string{"flower","flow","flight"}
  result := longestCommonPrefix(strs)
  fmt.Println(result)
}
