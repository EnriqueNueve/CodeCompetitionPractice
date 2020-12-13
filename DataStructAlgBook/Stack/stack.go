package stack

import (
	"errors"
)

// FILO: first in, last out
type Stack struct {
	data []int
}

func (s *Stack) push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() error {
	if len(s.data) > 0 {
		s.data = s.data[:len(s.data)-1]
		return nil
	} else {
		return errors.New("Stack is empty. No value to pop.")
	}
}

func (s *Stack) top() (int,error) {
  if len(s.data) > 0 {
    return s.data[len(s.data)-1], nil
  } else {
    return -1, errors.New("Stack is empty. No top value.")
  }
}
