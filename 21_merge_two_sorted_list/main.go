package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//選擇較小的節點做為下一個節點，直到兩個LIST維空
//也可用RECURSIVE解
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	newhead := &ListNode{}
	now := newhead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			now.Next = l2
			break
		}
		if l2 == nil {
			now.Next = l1
			break
		}
		if l1.Val < l2.Val {
			now.Next = l1
			now = now.Next
			l1 = l1.Next
		} else {
			now.Next = l2
			now = now.Next
			l2 = l2.Next
		}
	}
	return newhead.Next
}

func main() {
	fmt.Println("hello world")
}
