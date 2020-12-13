package stack

import (
    "testing"
)

func Testpush(t *testing.T) {
  s := Stack{data: []int{1,2,3}}
  s.push(4)
  if s.data[3] != 4 {
    t.Errorf("Pushed value to stack failed!")
  }
}

func Testpop(t *testing.T) {
  s := Stack{data: []int{1,2,3}}
  err := s.pop()
  if s.data[2] != 3 {
    t.Errorf("Pop failed to remove top value!")
  }

  s = Stack{}
  err = s.pop()
  if err == nil {
    t.Errorf("Popped empty list. Failed to produce error!")
  }
}

func Testtop(t *testing.T) {
  s := Stack{data: []int{1,2,3}}
  top_val, top_error := s.top()
  if top_val != 3 {
    t.Errorf("Pop failed to remove top value!")
  }

  s = Stack{}
  top_val, top_error = s.top()
  if top_error == nil {
    t.Errorf("Pop failed to identify empty stack!")
  }
}
