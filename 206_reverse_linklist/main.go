package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/*recursion 版本*/
func reverseListR(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	tail := head.Next
	newh := reverseListR(head.Next)
	tail.Next = head
	head.Next = nil
	return newh
}

/*iteratively 版本 記錄所有節點在一路反向建立*/
func reverseListI(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	allnode := []*ListNode{}
	for ;head!=nil;head = head.Next{
		allnode = append(allnode,head)
	}
	var newh *ListNode
	newh = allnode[len(allnode)-1]
	tmp := newh
	for i:=len(allnode)-2;i>=0;i--{
		tmp.Next = allnode[i]
		tmp = tmp.Next
	}
	tmp.Next = nil
	return newh
}

/*iteratively 版本 一路反向建立回去*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head.Next
	head.Next = nil
	newh := head
	now := head

	for ;tmp!=nil;{
		newh = tmp
		tmp = tmp.Next
		newh.Next = now
		now = newh
	}
	
	return newh
}



func main()  {
	fmt.Println("hello world")
}