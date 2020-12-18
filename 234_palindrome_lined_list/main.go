package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
	
func isPalindrome(head *ListNode) bool {
	if head==nil||head.Next == nil{
		return true
	}
	cnt,cur := 0,head
	for ;cur!=nil;cur=cur.Next{
		cnt++
	}
	first,second := head,head
	if cnt%2==0{
		step := (cnt-2)/2
		cur,n := head,head.Next
		for i:=0;i<step;i++{
			tmp := n.Next
			n.Next = cur
			cur = n
			n = tmp
		}
		head.Next = nil
		first,second = cur,n
	}else{
		step := (cnt-3)/2
		cur,n := head,head.Next
		for i:=0;i<step;i++{
			tmp := n.Next
			n.Next = cur
			cur = n
			n = tmp
		}
		head.Next = nil
		first,second = cur,n.Next
	}
	for first!=nil && second!=nil{
		if first.Val != second.Val{
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func main() {
	fmt.Println("hello world")
}
