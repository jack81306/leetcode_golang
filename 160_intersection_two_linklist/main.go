package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cura,curb := headA,headB
	for cura!=curb{
		if cura==nil{
			cura = headB
		}else{
			cura = cura.Next
		}
		if curb==nil{
			curb = headA
		}else{
			curb = curb.Next
		}
	}
	return cura
}

func main() {
	fmt.Println("hello world")
}