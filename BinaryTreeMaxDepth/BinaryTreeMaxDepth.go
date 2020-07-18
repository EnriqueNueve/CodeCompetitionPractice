/*
Binary Tree

* Binary Tree has at max two nodes
* Value on left is less than right value
* Root: level zero of the tree
* Child: each node on the tree that's not the root
* Internal node: each node with at least one child
* Leaf: each node that has no children
* Subtree: the set of node with a certain node as root

*/


package main


import (
        "fmt"
)


type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}


func maxDepth(root *TreeNode) int {
        if root == nil {
                return 0
        } else if root.Left == nil && root.Right == nil {
                return 1
        } else {
                var l int = maxDepth(root.Left)
                var r int = maxDepth(root.Right)
                if l > r {
                        return 1 + l
                } else {
                        return 1 + r
                }
        }
}


func main() {
        // [3,9,20,null,null,15,7]
        root := TreeNode{Val: 3}
        l1_node := TreeNode{Val: 9, Left: nil, Right: nil}
        r1_node := TreeNode{Val: 20}
        root.Right = &l1_node
        root.Left = &r1_node
        r1_l2 := TreeNode{Val: 15}
        r1_r2 := TreeNode{Val: 7}
        r1_node.Right = &r1_r2
        r1_node.Left = &r1_l2

        fmt.Println(maxDepth(&root))
}
