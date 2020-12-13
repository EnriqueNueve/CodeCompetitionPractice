package queue

import (
	"testing"
)

func Testpush(t *testing.T) {
	q := Queue{data: []int{1, 2, 3}}
	q.push(4)
	if q.data[3] != 4 {
		t.Errorf("Pushed value to queue failed!")
	}
}

func Testpop(t *testing.T) {
	q := Queue{data: []int{1, 2, 3}}
	_, err := q.pop()
	if q.data[1] != 2 {
		t.Errorf("Pop failed to remove top value!")
	}

	q = Queue{}
	_, err = q.pop()
	if err == nil {
		t.Errorf("Popped empty list. Failed to produce error!")
	}
}

func Testtop(t *testing.T) {
	q := Queue{data: []int{1, 2, 3}}
	top_val, top_error := q.top()
	if top_val != 1 {
		t.Errorf("Pop failed to remove top value!")
	}

	q = Queue{}
	top_val, top_error = q.top()
	if top_error == nil {
		t.Errorf("Pop failed to identify empty Queue!")
	}
}
