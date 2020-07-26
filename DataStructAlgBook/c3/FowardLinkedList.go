/*
____ Notes ____

Linked List: two-cells, in each oof which the first cell contains
            a pointer to a list element and the second cell
            contains a pointer to either the empty list or another cell.

*/

package main

import (
	"fmt"
)

// Struct for nodes of LinkedList
type LinkedListNode struct {
	Next *LinkedListNode
	Val  int
}

func main() {
	fmt.Println("Lists, Recursion, Stacks, Queues" + "\n")

	// Declare LinkedList
	head := LinkedListNode{Val: 2}
	n1 := LinkedListNode{Val: 3}
	head.Next = &n1
	n2 := LinkedListNode{Val: 4}
	n1.Next = &n2
	PrintLinkedList(&head)
	fmt.Println("\n")

	// Insert value to front of LinkedList
	head.MakeList(1)
	PrintLinkedList(&head)
	fmt.Println("\n")

	// Replace value in linked list
	head.ReplaceVal(3, 100)
	PrintLinkedList(&head)
	fmt.Println("\n")

	// Get length of linkes list 
	len_LinkedList := GetLen(&head)
	fmt.Println(len_LinkedList)

}
