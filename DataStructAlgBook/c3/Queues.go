/*
Queues: ideal data structures to model a First-In-First-Out (FIFO)
*/

package main

import (
	"fmt"
)

type Queue []int

func main() {
	var queue Queue

	queue.Push(1) // Add values to queue
	queue.Push(2)
	queue.Push(3)

	// When pop, first in was 1 thus, 1 one will come out
	fmt.Println(queue.Pop())
	// Then 2
	fmt.Println(queue.Pop())
	// Then 3
	fmt.Println(queue.Pop())

}
