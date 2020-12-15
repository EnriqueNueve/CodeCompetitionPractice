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
	Val   int
	Left  *Node
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
		return (1 + size(left(tree)) + size(right(tree)))
	}
}

func height(tree *Node) int {
	return int(math.Ceil(math.Log2(float64(size(tree))+1) - 1))
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

func isIn(tree *Node, val int) bool {
	if isEmpty(tree) == true {
		return false
	} else if val == root(tree) {
		return true
	} else if val < root(tree) {
		return isIn(left(tree), val)
	} else {
		return isIn(right(tree), val)
	}
}

func delete(tree *Node, val int) *Node {
  if isEmpty(tree) == true {
    fmt.Println("Value is not in tree.")
    return tree
  } else {
    if val < root(tree) {
      return MakeTree(root(tree), delete(left(tree),val),right(tree))
    } else if val > root(tree) {
      return MakeTree(root(tree), left(tree), delete(right(tree),val))
    } else {
      if isEmpty(left(tree)) {
        return right(tree)
      } else if isEmpty(right(tree)) {
        return left(tree)
      } else {
        return MakeTree(smallestNode(right(tree)),left(tree),removeSmallestNode(right(tree)))
      }
    }
  }
}

func smallestNode(tree *Node) int {
  if isEmpty(left(tree)) {
    return root(tree)
  } else {
    return smallestNode(left(tree))
  }
}

func removeSmallestNode(tree *Node) *Node {
  if isEmpty(left(tree)) == true {
    return right(tree)
  } else {
    return MakeTree(root(tree),removeSmallestNode(left(tree)),right(tree))
  }
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

func (t *Node) TraverseBalance(q *Queue) {
	if t.Left != nil {
		t.Left.TraverseBalance(q)
	}
	q.push(t.Val)
	if t.Right != nil {
		t.Right.TraverseBalance(q)
	}
}

func BuildTree(data []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	new_tree := Leaf(data[mid])
	new_tree.Left = BuildTree(data, start, mid-1)
	new_tree.Right = BuildTree(data, mid+1, end)
	return new_tree
}

type Queue struct {
	data []int
}

// Adds value to Queue
func (q *Queue) push(val int) {
	q.data = append(q.data, val)
}

func main() {
	tree := MakeTree(8, MakeTree(3, Leaf(1), MakeTree(6, EmptyTree(), Leaf(7))), MakeTree(11, MakeTree(9, EmptyTree(), Leaf(10)), MakeTree(14, Leaf(12), Leaf(15))))
	fmt.Println("Is the tree empty: ", isEmpty(tree))
	fmt.Println("Root value of the tree: ", root(tree))
	printTree(tree, height(tree))
	fmt.Println("Size of tree: ", size(tree))
	fmt.Println("Height of tree: ", height(tree))
	tree = tree.BalanceTree()
	printTree(tree, height(tree))
	fmt.Println("------------------------------\n")

	btree := MakeTree(10, EmptyTree(), EmptyTree())
	btree = insert(11, btree)
	btree = insert(12, btree)
	btree = insert(9, btree)
	printTree(btree, height(btree))
	fmt.Println("Is 9 in the tree: ", isIn(btree, 9))
	fmt.Println("Is 1 in the tree: ", isIn(btree, 1))
  btree = delete(btree,9)
  printTree(btree, height(btree))
}
