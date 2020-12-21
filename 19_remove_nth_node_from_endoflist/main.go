package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newhead := new(ListNode)
	newhead.Next = head
	slow,fast := newhead,head
	for i:=0;i<n;i++{
		fast = fast.Next
	}
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	slow.Next = slow.Next.Next
	return newhead.Next
}

func main() {
	fmt.Println("hello world")
}
