package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head==nil{
		return nil
	}
	one,two := head,head
	for one!=nil&&two!=nil{
		two = two.Next
		if two == nil{
			return nil
		}
		two = two.Next
		one = one.Next
		if one==two{
			break
		}
	}
    if one==nil||two==nil{
        return nil
    }
	one = head
	for one!=two{
		one = one.Next
		two = two.Next
	}
	return one	
}




func main() {
	fmt.Println("hello world")
}
