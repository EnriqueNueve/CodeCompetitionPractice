package linkedlist

import (
	"testing"
)

func Testinsert(t *testing.T) {
	ll := MakeList(3, MakeList(2, MakeList(1, &LinkedListNode{})))
	ll.insert(4)

	if ll.Next.Next.Next.Next.Val != 4 {
		t.Errorf("Failed to insert value into linked list!")
	}
}

func TestisEmpty(t *testing.T) {
	ll := &LinkedListNode{}
	if ll.isEmpty() == false {
		t.Errorf("Failed to identify empty linked list!")
	}
}

func TestisIn(t *testing.T) {
	ll := MakeList(3, MakeList(2, MakeList(1, &LinkedListNode{})))
	if isIn(ll, 3) != true {
		t.Errorf("Failed to check if value is in linked list!")
	}

	if isIn(ll, 10) != false {
		t.Errorf("Failed to check if value is in linked list!")
	}
}

func Testreverse(t *testing.T) {
	ll := MakeList(3, MakeList(2, MakeList(1, &LinkedListNode{})))
  ll = ll.reverse()

	if ll.Val != 1 {
		t.Errorf("Failed to reverse linked list!")
	}
	if ll.Next.Val != 2 {
		t.Errorf("Failed to reverse linked list!")
	}
  if ll.Next.Next.Val != 3 {
    t.Errorf("Failed to reverse linked list!")
  }
}
