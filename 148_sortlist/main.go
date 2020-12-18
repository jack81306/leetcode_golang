package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
    if head.Next==nil{
        return head
    }
	cnt := 0
	cur := head
	for cur!=nil{
		cur = cur.Next
		cnt++
	}
	l := head
	cur = head
	for i:=0;i<int(cnt/2)-1;i++{
		cur = cur.Next
	}
	r := cur.Next
	cur.Next = nil
	l = sortList(l)
	r = sortList(r)
	newhead := new(ListNode)
	cur = newhead
	for l!=nil||r!=nil{
		if l==nil{
			cur.Next = r
			r = nil
			break
		}
		if r==nil{
			cur.Next = l
			l = nil
			break
		}
		if l.Val > r.Val{
			cur.Next = r
			r = r.Next
		}else{
			cur.Next = l
			l = l.Next
		}
		cur = cur.Next
	}

	return newhead.Next
}


func printls(head *ListNode){
	for head!=nil{
		fmt.Print(head.Val," ")
		head = head.Next
	}
	fmt.Println("")
}

func main() {
	fmt.Println("hello world")
}
