package main

import (
	"fmt"
)
type Node struct {
    Val int
    Next *Node
    Random *Node
}

   
   
func copyRandomList(head *Node) *Node {
	if head==nil{
		return nil
	}
	cur := head
	for cur!=nil{
		n := new(Node)
		n.Val = cur.Val
		n.Next = cur.Next
		cur.Next = n
		cur = cur.Next.Next
	}
	cur = head
	cop := head.Next
	for cur != nil{
		if cur.Random == nil{
			cop.Random = nil
		}else{
			cop.Random = cur.Random.Next
		}
		if cop.Next != nil{
			cop = cop.Next.Next
			cur = cur.Next.Next
		}else{
			cur = nil
		}
	}

	newhead,oldhead := new(Node),new(Node)
	tailnew,tailold := newhead,oldhead
	cur = head
	count := 0
	for cur!=nil{
		if count%2==0{
			tailold.Next = cur
            tailold = tailold.Next
		}else{
			tailnew.Next = cur
            tailnew = tailnew.Next
		}
		count++
		cur = cur.Next
	}
    tailold.Next = nil
    
	return newhead.Next
}

func main() {
	fmt.Println("hello world")
}
