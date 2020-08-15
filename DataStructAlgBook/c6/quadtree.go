/*

Notes:

* Trees consists of an ordered set of linked nodes
  in a connected graph, in which each node has at
  most one parent node, and zero or more children
  nodes with a specific order.

* Tree as consisting of nodes (also called vertices or points)
  and edges (also called lines, or, in order to stress the
  directedness, arcs) with a tree-like structure.

* Refer to the label of a node as its value.

* Root: ‘top level’ node known as the root.

* Nodes that have the same parent are known as
  siblings – siblings are, by definition, always
  on the same level.

* The size of a tree is given by the number of nodes it contains.

* The maximal length of a path in a tree is also called the height of the tree.

* Quadtree: is a particular type of tree in which each leaf-node
  is labelled by a value and each non-leaf node has exactly four children.

*/

package main

import (
	"fmt"
)

type QuadTreeNode struct {
	Val  int
	luqt *QuadTreeNode
	ruqt *QuadTreeNode
	llqt *QuadTreeNode
	rlqt *QuadTreeNode
}

func main() {
	node_A := QuadTreeNode{Val:2}
	node_B := QuadTreeNode{Val:3}
	node_C := QuadTreeNode{Val:4}
	node_D := QuadTreeNode{Val:5}
	root := QuadTreeNode{Val: 1,luqt:&node_A,ruqt:&node_B,llqt:&node_C,rlqt:&node_D}
	fmt.Println("Root value: ",root.Val)
	fmt.Println("Node A value: ",root.luqt.Val)
	fmt.Println("Node B value: ",root.ruqt.Val)
	fmt.Println("Node C value: ",root.llqt.Val)
	fmt.Println("Node D value: ",root.rlqt.Val)
}
