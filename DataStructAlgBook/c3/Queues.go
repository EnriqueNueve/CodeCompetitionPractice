/*
Queues: ideal data structures to model a First-In-First-Out (FIFO)
*/

package main

import (
	"fmt"
	"errors"
)

type Queue []int

// IsEmpty: check if Queue is empty
func isEmpty(q *Queue) bool {
	return len(*q) == 0
}

// Push a new value into the Queue
func push(element int, q *Queue) {
	*q = append(*q, element) // Simply append the new value to the Queue
}

// Returns top element of Queue without removing
func Top(q *Queue) int {
	return (*q)[0]
}

// Remove and return first element of Queue. Return false if Queue is empty.
func pop(q *Queue) (int, error) {
	if isEmpty(q) {
		return -1, errors.New("Error: queue is empty")
	} else {
		element := (*q)[0] // Get first element in slice
		*q = (*q)[1:]      // Remove it from the Queue by slicing it off.
		return element, nil
	}
}


func main() {
	var queue Queue

	push(1,&queue) // Add values to queue
	push(2,&queue)
	push(3,&queue)

	// When pop, first in was 1 thus, 1 one will come out
	fmt.Println(pop(&queue))
	// Then 2
	fmt.Println(pop(&queue))
	// Then 3
	fmt.Println(pop(&queue))
	// Pop and get error since empty
	fmt.Println(pop(&queue))

}
