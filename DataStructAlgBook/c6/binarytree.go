package main

import (
	"errors"
	"fmt"
)

// Node of tree
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Empty pointer for dead branch
type EmptyNode struct {
	E *BinaryTreeNode
}

// Dummy function to create dead branch
func EmptyTree() *BinaryTreeNode {
	temp := EmptyNode{}
	return temp.E
}

// Creates node with value and two pointers: left and right
func MakeTree(val int, left *BinaryTreeNode, right *BinaryTreeNode) *BinaryTreeNode {
	parent := BinaryTreeNode{Val: val, Left: left, Right: right}
	return &parent
}

// Place LeafNode on tree. LeafNode has a nil left and right pointer.
func Leaf(val int) *BinaryTreeNode {
	leaf := BinaryTreeNode{Val: val, Left: nil, Right: nil}
	return &leaf
}

// Get root node value
func root(tree *BinaryTreeNode) (int, error) {
	if tree == nil {
		return -1, errors.New("Root node value does not exist!")
	}
	return tree.Val, nil
}

// Get pointer of left node of relative parent node
func left(tree *BinaryTreeNode) *BinaryTreeNode {
	left_t := tree.Left
	if left_t == nil {
		fmt.Println("WARNING: left tree is empty!")
	}
	return left_t
}

// Get pointer of right node of relative partent node
func right(tree *BinaryTreeNode) *BinaryTreeNode {
	right_t := tree.Right
	if right_t == nil {
		fmt.Println("WARNING: right tree is empty!")
	}
	return right_t
}

// Check if tree pointer is nil
func isEmpty(tree *BinaryTreeNode) bool {
	if tree == nil {
		return true
	} else {
		return false
	}
}

// Count number of nodes in a tree
func size(tree *BinaryTreeNode) int {
	var counter int = 0
	if isEmpty(tree) {
		return 0
	} else {
		return (1 + size(left(tree)) + size(right(tree)))
	}
	return counter
}

func height(tree *BinaryTreeNode) int {
	if isEmpty(tree) {
		return 0
	}
	if isEmpty(left(tree)) == true && isEmpty(right(tree)) == true {
		return 1
	} else {
		return height(left(tree)) + height(right(tree))
	}
}

func main() {
	// MakeTree returns pointer to BinaryTreeNode struct
	//tree := MakeTree(1, MakeTree(2, EmptyTree(), Leaf(3)), MakeTree(4, Leaf(5), Leaf(6)))
  tree := MakeTree(1, EmptyTree(),EmptyTree())


	// test root()
	// test 1
	root_val, error := root(tree)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(root_val)

	// test 2
	tree_empty := EmptyTree()
	root_val, error = root(tree_empty)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Root value of the tree: ", root_val, "\n")

	// test left()
	fmt.Println(left(tree), "\n")

	// test size()
	fmt.Println("Size of tree: ", size(tree), "\n")

  // test height
  fmt.Println("Height of tree: ",height(tree), "\n")

	// test isEmpty()
	fmt.Println(isEmpty(EmptyTree()))
	fmt.Println(isEmpty(tree.Left))
	fmt.Println(isEmpty(tree.Right), "\n")
}
