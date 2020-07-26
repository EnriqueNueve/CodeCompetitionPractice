package main

// IsEmpty: check if Queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push a new value into the Queue
func (q *Queue) Push(element int) {
	*q = append(*q, element) // Simply append the new value to the Queue
}

// Returns top element of Queue without removing
func (q *Queue) Top() int {
	return (*q)[0]
}

// Remove and return first element of Queue. Return false if Queue is empty.
func (q *Queue) Pop() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	} else {
		element := (*q)[0] // Get first element in slice 
		*q = (*q)[1:]      // Remove it from the Queue by slicing it off.
		return element, true
	}
}
