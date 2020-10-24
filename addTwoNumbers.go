package main

import (
  "fmt"
)

type ListNode struct {
  Val int
  Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  head := &ListNode{}
  for prev, sum := head, 0; l1 != nil || l2 != nil || sum > 0; sum = sum/10 {
    if l1 != nil {
      sum += l1.Val
      l1 = l1.Next
    }
    if l2 != nil {
      sum += l2.Val
      l2 = l2.Next
    }
    prev.Next= &ListNode{Val:sum%10}
    prev = prev.Next
  }
  // return head.Next since first node is empty
  return head.Next
}

func main(){
  // 243
  l1 := ListNode{Val:2}
  l1.Next = &ListNode{Val:4}
  l1.Next.Next = &ListNode{Val:3}
  l1.Next.Next.Next = &ListNode{Val:3}

  //564
  l2 := ListNode{Val:5}
  l2.Next = &ListNode{Val:6}
  l2.Next.Next = &ListNode{Val:4}

  ll := addTwoNumbers(&l1,&l2)
  fmt.Println((*ll).Val,(*ll).Next.Val,(*ll).Next.Next.Val,(*ll).Next.Next.Next.Val)
}
