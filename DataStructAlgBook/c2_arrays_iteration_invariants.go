/*
____ Notes ____

~ Data is stored as a pattern of bits but we interact with
    higher level objects like ints, char, etc.

Array: ordered collection of items.

Loop: method to repeat a process a number of times.

Invariant: something that does not change during execution of a program.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays, Iteration, Invariants!")

  // Array
  a := [5]int{1,2,3,4,5}

  fmt.Println("\n")

  // Loop
  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }

}
