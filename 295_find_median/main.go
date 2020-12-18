package main

import (
	"container/heap"
	"fmt"
)

type maxheap []int
type minheap []int

func (h maxheap) Len()int {return len(h)}
func (h maxheap) Less(i,j int) bool {return h[i]>h[j]}
func (h maxheap) Swap(i,j int){h[i],h[j] = h[j],h[i]}
func (h *maxheap) Push(x interface{}){
	*h = append(*h,x.(int))
}
func (h *maxheap) Pop() interface{}{
	oldh := *h
	res := oldh[len(oldh)-1]
	*h = oldh[:len(oldh)-1]
	return res
}

func (h minheap) Len()int {return len(h)}
func (h minheap) Less(i,j int) bool {return h[i]<h[j]}
func (h minheap) Swap(i,j int){h[i],h[j] = h[j],h[i]}
func (h *minheap) Push(x interface{}){
	
	*h = append(*h,x.(int))
}
func (h *minheap) Pop() interface{}{
	oldh := *h
	res := oldh[len(oldh)-1]
	*h = oldh[:len(oldh)-1]
	return res
}



type MedianFinder struct {
	maxh *maxheap
	minh *minheap
}




/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{&maxheap{},&minheap{}}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.maxh.Len() == 0{
		heap.Push(this.maxh,num)
		return
	}
	if num > (*this.maxh)[0]{
		heap.Push(this.minh,num)
	}else{
		heap.Push(this.maxh,num)
	}
	this.modify()
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxh.Len() == this.minh.Len(){
		return float64((*this.maxh)[0]+(*this.minh)[0])/2
	}
	return float64((*this.maxh)[0])
}

func (this *MedianFinder)modify()  {
	if this.maxh.Len() < this.minh.Len(){
		m := heap.Pop(this.minh)
		heap.Push(this.maxh,m)
	}else if  this.maxh.Len()-this.minh.Len()>1{
		m := heap.Pop(this.maxh)
		heap.Push(this.minh,m)
	}
}


func main() {
	fmt.Println("hello world")
}
