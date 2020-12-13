// TODO: Implement check if balanced -> bool

package main

import (
	"fmt"
)

// ------
// Binary Tree functions, methods, and structs
// ------

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func MakeBTree(data int) *Node {
	return &Node{Left: nil, Right: nil, Data: data}
}

func MakeBTreeLR(data int, Left *Node, Right *Node) *Node {
	return &Node{Left: Left, Right: Right, Data: data}
}

func right(t *Node) *Node {
	if t.Right != nil {
		return t.Right
	} else {
		return nil
	}
}

func left(t *Node) *Node {
	if t.Left != nil {
		return t.Left
	} else {
		return nil
	}
}

func (t *Node) Insert(data int) *Node {
	if data < t.Data {
		if t.Left != nil {
			t.Left.Insert(data)
		} else {
			t.Left = MakeBTree(data)
		}
	} else {
		if t.Right != nil {
			t.Right.Insert(data)
		} else {
			t.Right = MakeBTree(data)
		}
	}
	return t
}

func (t *Node) Contains(data int) bool {
	// Check root
	if t.Data == data {
		return true
	}
	// Check left
	if data < t.Data {
		if t.Left != nil {
			t.Left.Contains(data)
		} else {
			return false
		}
	} else { //Check right
		if t.Right != nil {
			t.Right.Contains(data)
		} else {
			return false
		}
	}
	return false
}

func printTree(t *Node, level int) {
	if t != nil {
		// this is used for spacing
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---<["
		level++

		// recursive call to traverse tree
		printTree(t.Right, level)
		fmt.Printf(format+"%d]\n", t.Data)
		printTree(t.Left, level)
	}
}

func (t *Node) TraverseBalance(q *Queue) {
	if t.Left != nil {
		t.Left.TraverseBalance(q)
	}
	q.push(t.Data)
	if t.Right != nil {
		t.Right.TraverseBalance(q)
	}
}

func BuildTree(data []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	new_tree := MakeBTree(data[mid])
	new_tree.Left = BuildTree(data, start, mid-1)
	new_tree.Right = BuildTree(data, mid+1, end)
	return new_tree
}

// Create new  Balanced Binary Search Tree
func (t *Node) BalanceTree() *Node {
	// Place all elements of Binary Tree in Queue
	q := Queue{}
	t.TraverseBalance(&q)

	// Build new tree
	new_tree := BuildTree(q.data, 0, len(q.data)-1)

	return new_tree
}

// ------
// Queue functions, methods, and structs
// ------

type Queue struct {
	data []int
}

// Adds value to Queue
func (q *Queue) push(val int) {
	q.data = append(q.data, val)
}

// ------
// Main function
// ------

func main() {
	//tree := MakeBTree(4)
	//tree.Insert(5).Insert(6).Insert(7)
	//tree.Insert(3).Insert(2).Insert(1)

	//tree := MakeBTree(30)
	//tree.Insert(20).Insert(10)

	//printTree(tree,3)
	//tree = tree.BalanceTree()
	//fmt.Println()
	//printTree(tree,3)

	tree := MakeBTree(10)
	tree.Insert(5).Insert(6).Insert(4).Insert(12).Insert(13).Insert(11)

	printTree(tree,3)
	printTree(right(tree),3)
	printTree(left(tree),3)
}
