/*
Functions and methods to be called on LinkedListNode struct
*/

package main

import "fmt"

// Returns int of length of list
func GetLen(head *LinkedListNode) int {
	var counter int = 1
	for head.Next != nil {
		head = head.Next
		counter++
	}
	return counter
}

// Insets LinkedListNode to front of LinkedList
// this func is a 'method' which allows changes to structs
func (head *LinkedListNode) MakeList(element int) {
	newHead := LinkedListNode{Val: head.Val, Next: head.Next}
	head.Next = &newHead // set head.Next to new node that points to rest of list
	head.Val = element   // set head.Val to new value
}

// Insets LinkedListNode to front of LinkedList
// this func is a 'method' which allows changes to structs
func (head *LinkedListNode) ReplaceVal(index int, element int) {
	header := head
	for i := 0; i < index; i++ {
		header = header.Next
	}
	header.Val = element
}

// Print elements of LinkedList
func PrintLinkedList(node *LinkedListNode) {
	fmt.Println(node.Val)
	if node.Next != nil {
		PrintLinkedList(node.Next)
	}
}
