package main

import (
	"fmt"
)

// const to turn on and off left and right method warning for empty
const panicEmptyError = false

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

// Get root node value
func root(tree *BinaryTreeNode) int {
	if tree == nil {
		return -1
	}
	return tree.Val
}

// prints out a visual of the tree
// function obtained from: https://flaviocopes.com/
// golang-data-structure-binary-search-tree/
func printTree(tree *BinaryTreeNode, level int) {
	if tree != nil {
		// this is used for spacing
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---<["
		level++

		// recursive call to traverse tree
		printTree(tree.Right, level)
		fmt.Printf(format+"%d]\n", tree.Val)
		printTree(tree.Left, level)
	}
}

// Check if tree pointer is empty
func isEmpty(tree *BinaryTreeNode) bool {
	// Check if empty pointer from EmptyTree()
	if tree == nil {
		return true
	} else {
		return false
	}
}

// Get pointer of left node of relative parent node
func left(tree *BinaryTreeNode) *BinaryTreeNode {
	left_t := tree.Left // get left pointer from parent node
	if left_t == nil && panicEmptyError == true {
		fmt.Println("WARNING: left tree is empty!")
	}
	return left_t
}

// Get pointer of right node of relative parent node
func right(tree *BinaryTreeNode) *BinaryTreeNode {
	right_t := tree.Right // get right pointer from parent node
	if right_t == nil && panicEmptyError == true {
		fmt.Println("WARNING: right tree is empty!")
	}
	return right_t
}

// Dummy function to create dead branch.
func EmptyTree() *BinaryTreeNode {
	temp := EmptyNode{}
	return temp.E
}

// Creates node with value and two pointers- left and right.
func MakeTree(val int, left *BinaryTreeNode, right *BinaryTreeNode) *BinaryTreeNode {
	parent := BinaryTreeNode{Val: val, Left: left, Right: right}
	return &parent
}

// Creates node with value and two pointers- left and right.
func insert(val int, tree *BinaryTreeNode) *BinaryTreeNode {
	if isEmpty(tree) == true {
		return MakeTree(val, EmptyTree(), EmptyTree())
	} else if val < root(tree) {
		return MakeTree(root(tree), insert(val, left(tree)), right(tree))
	} else if val > root(tree) {
		return MakeTree(root(tree), left(tree), insert(val, right(tree)))
	} else {
		print("Error: violated assumption in procedure insert.")
		return EmptyTree()
	}
}

// Checks if value is in tree
func isIn(val int, tree *BinaryTreeNode) bool {
	if isEmpty(tree) {
		return false
	} else if val == root(tree) {
		return true
	} else if val < root(tree) {
		return isIn(val, left(tree))
	} else {
		return isIn(val, right(tree))
	}
}

func delete(val int, tree *BinaryTreeNode) *BinaryTreeNode {
	if isEmpty(tree) {
		fmt.Println("Tree is empty, no node to delete.")
		return tree
	} else {
		if val < root(tree) {
			return MakeTree(root(tree), delete(val, left(tree)), right(tree))
		} else if val > root(tree) {
			return MakeTree(root(tree), left(tree), delete(val, right(tree)))
		} else {
			if isEmpty(left(tree)) {
				return right(tree)
			} else if isEmpty(right(tree)) {
				return left(tree)
			} else {
				return MakeTree(smallestNode(right(tree)), left(tree), removeSmallestNode(right(tree)))
			}
		}
	}
}

func smallestNode(tree *BinaryTreeNode) int {
	if isEmpty(left(tree)) {
		return root(tree)
	}
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree.Val
}

func removeSmallestNode(tree *BinaryTreeNode) *BinaryTreeNode {
	if isEmpty(left(tree)) {
		return right(tree)
	} else {
		return MakeTree(root(tree), removeSmallestNode(left(tree)), right(tree))
	}
}

func main() {

	tree := MakeTree(10, EmptyTree(), EmptyTree())
	tree = insert(11, insert(13, insert(6, insert(12, insert(8, tree)))))

	// one branch down
	fmt.Println(tree.Left.Val)
	fmt.Println(tree.Right.Val)
	fmt.Println()

	// two branch down
	fmt.Println(tree.Left.Left.Val)
	fmt.Println(tree.Right.Right.Val)
	fmt.Println(tree.Right.Left.Val)
	fmt.Println()

	// check if value is in tree
	fmt.Println("Is 6 in the tree? ", isIn(6, tree))
	fmt.Println("Is 100 in the tree? ", isIn(100, tree))
	fmt.Println()

	// view tree
	printTree(tree, 4)
	fmt.Println("------------------------------------------------------")

	// delete node
	fmt.Println("Removing node with value 10.")
	tree = delete(10, tree)
	printTree(tree, 4)
	fmt.Println()

}
