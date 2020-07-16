package main


import (
        "fmt"
)


type ListNode struct {
        Val  int
        Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
        current := head
        var next *ListNode = nil
        var prev *ListNode = nil
        for current != nil {
          next = current.Next
          current.Next = prev
          prev = current
          current = next
        }
        return prev
}


func PrintLinkedList(node *ListNode) {
        fmt.Println(node.Val)
        if node.Next != nil {
                PrintLinkedList(node.Next)
        }
}


func main() {
        node_1 := ListNode{Val: 1}

        node_2 := ListNode{Val: 2}
        node_1.Next = &node_2

        node_3 := ListNode{Val: 3}
        node_2.Next = &node_3

        results := reverseList(&node_1) // pass pointer to first node
        PrintLinkedList(results)

}
