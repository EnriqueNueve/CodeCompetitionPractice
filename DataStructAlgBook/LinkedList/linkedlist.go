package main

/*
LinkedListNode, struct
MakeList, function
isEmtpy, method
getLen, function
insert, method
printLinkedList, function
*/

import (
	"fmt"
)

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

func MakeList(element int, node *LinkedListNode) *LinkedListNode {
	if (LinkedListNode{}) == (*node) {
		node.Val = element
		return node
	} else {
		newHead := LinkedListNode{Val: node.Val, Next: node.Next}
		node.Next = &newHead
		node.Val = element
		return node
	}
}

func (node *LinkedListNode) isEmpty() bool {
	if (LinkedListNode{}) == (*node) {
		return true
	} else {
		return false
	}
}

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

// Add value at the end of the LinkedList
func (node *LinkedListNode) insert(val int) *LinkedListNode {
	head := node
	new_node := LinkedListNode{Val: val, Next: nil}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &new_node
	return head
}

// Method to view all values of LinkedList
func printLinkedList(node *LinkedListNode){
	for node != nil {
		if node.Next != nil {
			fmt.Print("[",node.Val,"] -> ")
		} else {
			fmt.Print("[",node.Val,"]")
		}
		node = node.Next
	}
	fmt.Println("")
}

func (node *LinkedListNode) reverse() *LinkedListNode {
  length_ll := getLen(node)

  if length_ll > 1 {
    temp_pointerA := node
    temp_pointerB := new(LinkedListNode)
    head := node

    node = node.Next
    for i := 0; i < length_ll -1; i++ {
      temp_pointerB = node.Next
      node.Next = temp_pointerA
      temp_pointerA = node
      if i < length_ll - 2 {
        node = temp_pointerB
      }
    }
    head.Next = nil
    return node
  } else {
    return node
  }
}

func main() {
	ll := MakeList(3, MakeList(2, MakeList(1, &LinkedListNode{})))
	ll.insert(4).insert(5)
	printLinkedList(ll)
	fmt.Println("Len of linked list: ", getLen(ll))
	fmt.Println("Is linked list empty: ", ll.isEmpty())
	ll = ll.reverse()
	printLinkedList(ll)
}
