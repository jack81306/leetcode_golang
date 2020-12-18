package main

import (
	"fmt"
	"container/heap"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type minheap []*ListNode

func (h minheap) Len() int {return len(h)}
func (h minheap) Less(i, j int) bool { return (*h[i]).Val < (*h[j]).Val }
func (h minheap) Swap(i, j int)      { (*h[i]),(*h[j]) =(*h[j]), (*h[i]) }
func (h *minheap)Push(x interface{}){
	*h = append(*h,x.(*ListNode))
}
func (h *minheap)Pop()interface{}{
	old := *h
	res := old[len(old)-1]
	*h = old[:len(old)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
    if allnil(lists){
		return nil
	}
	newhead := new(ListNode)
	cur := newhead
	h := &minheap{}
	for _,v:=range lists{
		if v!=nil{
			heap.Push(h,v)
		}
	}
	for h.Len()!=0{
		min := heap.Pop(h)
		cur.Next = min.(*ListNode)
		cur = cur.Next
		if min.(*ListNode).Next!=nil{
			heap.Push(h,min.(*ListNode).Next)
		}
	}
	return newhead.Next
}

func allnil(lists []*ListNode)bool{
	for _,v := range lists{
		if v!=nil{
			return false
		}
	}
	return true
}


func main() {
	fmt.Println("hello world")
}
