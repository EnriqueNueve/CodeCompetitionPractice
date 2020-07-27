/*
Stack: ideal data structures to model a First-In-Last-Out (FILO),
        or Last-In-First-Out (LIFO)
*/

package main

import (
  "fmt"
  "errors"
)

// A type is made for a list called Stack. This allows us to
//  create methods on Stack, just like the methods of a class
type Stack []int

// Returns top element of stack without removing
func top(s *Stack) int {
  return (*s)[len(*s) - 1 ]
}

// Push a new value onto the stack
func Push(element int, s *Stack) {
	*s = append(*s, element) // Simply append the new value to the end of the stack
}

// IsEmpty: check if stack is empty
func isEmpty(s *Stack) bool {
	return len(*s) == 0
}

// Returns top element of stack without removing
func Top(s *Stack) (int,error) {
  if len(*s) == 0 {
    return -1, errors.New("Stack is empty, no top present")
  }
  return (*s)[len(*s) - 1 ], nil
}

// Remove and return top element of stack. Return false if stack is empty.
func pop(s *Stack) (int, error) {
	if isEmpty(s) {
		return -1, errors.New("Stack is empty, no value to pop")
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, nil
	}
}

// Pop all element of the stack
func Dump(s *Stack) {
  for len(*s) != 0 {
    pop(s)
  }
}

func main(){
  // create a stack variable of type Stack
  var stack Stack

  // add values to stack
  Push(1,&stack)
  Push(2,&stack)
  Push(3,&stack)
  Push(4,&stack)

  // check if stack is empty or not
  fmt.Println(isEmpty(&stack))

  // view top value of stack
  top_val, top_err := Top(&stack)
  if top_err != nil {
    fmt.Println(top_err)
  } else {
    fmt.Println("Top value: ", top_val)
  }

  // pop top element of stack
  pop_val, pop_err := pop(&stack)
  if pop_err != nil {
    fmt.Println(pop_err)
  } else {
    fmt.Println("Pop value: ", pop_val)
  }

  // clear stack
  Dump(&stack)
  fmt.Println(stack)
}
