package queue

import (
	"errors"
)

// FIFO: first in, first out
type Queue struct {
	data []int
}

// Checks if Queue has any values inside
func (q *Queue) isEmpty() bool {
	if len(q.data) == 0 {
		return false
	} else {
		return true
	}
}

// Adds value to Queue
func (q *Queue) push(val int) {
	q.data = append(q.data, val)
}

// Shows next value to be removed from Queue
func (q *Queue) top() (int, error) {
	if len(q.data) > 0 {
		return q.data[0], nil
	} else {
		return -1, errors.New("Queue is empty, no top value!")
	}
}

// Removes value from Queue
func (q *Queue) pop() (int, error) {
	if len(q.data) > 0 {
		pop_val := q.data[0]
		q.data = q.data[1:]
		return pop_val, nil
	} else {
		return -1, errors.New("Queue is empty, no value to pop!")
	}
}
