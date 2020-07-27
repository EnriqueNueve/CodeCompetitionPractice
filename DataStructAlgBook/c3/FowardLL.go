/*
____ Notes ____

Linked List: two-cells, in each oof which the first cell contains
            a pointer to a list element and the second cell
            contains a pointer to either the empty list or another cell.

*/

package main

import (
	"errors"
	"fmt"
)

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

// Inductive approach to filling LinkedList
// Ex: head := MakeList(2,MakeList(1, &LinkedListNode{}))
func MakeList(element int, node *LinkedListNode) *LinkedListNode {
	if (LinkedListNode{}) == (*node) { // Special case for empty list
		node.Val = element
		return node
	} else {
		newHead := LinkedListNode{Val: node.Val, Next: node.Next}
		node.Next = &newHead // set head.Next to new node that points to rest of list
		node.Val = element   // set head.Val to new value
		return node
	}
}

// Method to check if LinkedList is empty
func (node *LinkedListNode) isEmpty() bool {
	if (LinkedListNode{}) == (*node) {
		return true
	} else {
		return false
	}
}

// Method to check to retrieve first value of LinkedList
func (node *LinkedListNode) first() int {
	if node.isEmpty() == false {
		return node.Val
	} else {
		return -1
	}
}

// Method to view all values but first of LinkedList
func (node *LinkedListNode) rest() []int {
	var values []int
	node = node.Next
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return values
}

// Method to replace first value of LinkedList
func (node *LinkedListNode) replaceFirst(element int) {
	node.Val = element
}

// Function to repalce all values but first of LinkedList
func (node *LinkedListNode) replaceRest(new_vals []int) error {
	l := getLen(node)
	if l > 1 && len(new_vals)+1 == l {
		node = node.Next
		for node != nil {
			node.Val = new_vals[0]
			node = node.Next
			new_vals = new_vals[1:]
		}
	} else {
		return errors.New("New values longer than rest of LinkedList")
	}
	return nil
}

// Function to get length of LinkedList
func getLen(node *LinkedListNode) int {
	var counter int = 0
	if node != nil {
		counter++
	} else {
		return 0
	}
	for node != nil {
		node = node.Next
		counter++
	}
	return counter - 1
}

func main() {

	// Declare LinkedList by induction
	fmt.Println("Declare LinkedList and view values")
	head := MakeList(3, MakeList(2, MakeList(1, &LinkedListNode{})))
	fmt.Println(head.Val)
	fmt.Println(head.Next.Val)
	fmt.Println("\n")

	// Checks for empty LinkedList
	fmt.Println("Check if LinkedList is empty")
	fmt.Println(head.isEmpty())
	fmt.Println("\n")

	// View first value of LinkedList
	fmt.Println("View first value of LinkedList")
	fmt.Println(head.first())
	fmt.Println("\n")

	// View all values but first of LinkedList
	fmt.Println("View all values of LinkedList")
	fmt.Println(head.rest())
	fmt.Println("\n")

	// Get length of LinkedList
	fmt.Println("Length of LinkedList")
	fmt.Println(getLen(head))
	fmt.Println("\n")

	// Replace first value of LinkedList
	fmt.Println("Replace first value of LinkedList")
	head.replaceFirst(9)
	fmt.Println(head.first())
	fmt.Println(head.rest())
	fmt.Println("\n")

	// Replace all values but first
	fmt.Println("Replace all values but first of LinkedList")
	new_vals := []int{8, 9}
	err := head.replaceRest(new_vals)
	if err == nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(head.first())
	fmt.Println(head.rest())
	fmt.Println("\n")

}
