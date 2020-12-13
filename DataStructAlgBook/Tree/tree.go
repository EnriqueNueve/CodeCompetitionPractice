/*
A binary search tree is a binary tree that is either empty or satisfies the following conditions:
• All values occurring in the left subtree are smaller than that of the root.
• All values occurring in the right subtree are larger than that of the root.
• The left and right subtrees are themselves binary search trees.
*/

package main

import (
  "fmt"
  "math"
)

type Node struct {
  Val int
  Left *Node
  Right *Node
}

func EmptyTree() *Node {
  return nil
}

func MakeTree(root int, left *Node, right *Node) *Node {
  return &Node{Val: root, Left: left, Right: right}
}

func Leaf(root int) *Node {
  return &Node{Val: root, Left: nil, Right: nil}
}

func isEmpty(tree *Node) bool {
  if tree == nil {
    return true
  } else {
    return false
  }
}

func root(tree *Node) int {
  return tree.Val
}

func left(tree *Node) *Node {
  return tree.Left
}

func right(tree *Node) *Node {
  return tree.Right
}

func size(tree *Node) int {
  if isEmpty(tree) == true {
    return 0
  } else {
    return (1+size(left(tree))+size(right(tree)))
  }
}

func height(tree *Node) int {
  return int(math.Ceil(math.Log2(float64(size(tree))+1)-1))
}

func printTree(tree *Node, level int) {
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

func insert(val int, tree *Node) *Node {
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

func main() {
  tree := MakeTree(8, MakeTree(3,Leaf(1),MakeTree(6,EmptyTree(),Leaf(7))),MakeTree(11,MakeTree(9,EmptyTree(),Leaf(10)),MakeTree(14,Leaf(12),Leaf(15))))
  fmt.Println("Is the tree empty: ",isEmpty(tree))
  fmt.Println("Root value of the tree: ",root(tree))
  printTree(tree,height(tree))
  fmt.Println("Size of tree: ",size(tree))
  fmt.Println("Height of tree: ",height(tree))
  fmt.Println("------------------------------\n")

  btree := MakeTree(10,EmptyTree(),EmptyTree())
  btree = insert(11,btree)
  btree = insert(12,btree)
  btree = insert(9,btree)
  printTree(btree,height(btree))

}
