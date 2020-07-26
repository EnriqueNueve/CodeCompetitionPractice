/*
Stack: ideal data structures to model a First-In-Last-Out (FILO),
        or Last-In-First-Out (LIFO)
*/

package main

import (
  "fmt"
)

// A type is made for a list called Stack. This allows us to
//  create methods on Stack, just like the methods of a class
type Stack []int

func main(){
  var stack Stack // create a stack variable of type Stack

  stack.Push(1) // add values to stack
  stack.Push(2)
  stack.Push(3)
  stack.Push(4)

  top_val := stack.Top() // view top element of stack
  fmt.Println(top_val)

  stack.Dump() // Pop all elements of off stack
}
