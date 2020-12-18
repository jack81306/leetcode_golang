package main

import (
	"fmt"
)

type MinStack struct {
	s,s2 []int
	min int
}
/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{},[]int{},int(uint(1)<<31)}
}


func (this *MinStack) Push(x int)  {
	this.s = append(this.s,x)
	if x < this.min {
		this.min = x
	}
	this.s2 = append(this.s2,this.min)
}

func (this *MinStack) Pop()  {
	this.s = this.s[:len(this.s)-1]
	this.s2 = this.s2[:len(this.s2)-1]
    if len(this.s2) > 0{
        this.min = this.s2[len(this.s2)-1]
    }else{
        this.min = int(uint(1)<<31)
    }
}



func (this *MinStack) Top() int {
    return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}




func main() {
	fmt.Println("hello world")
}
