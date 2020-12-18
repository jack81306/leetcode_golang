package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head==nil{
		return false
	}
	one,two := head,head
	for one!=nil&&two!=nil{
		two = two.Next
		if two == nil{
			return false
		}
		two = two.Next
		one = one.Next
		if one==two{
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("hello world")
}
