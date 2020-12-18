package main

import (
	"fmt"
)

type deque []int

func (d *deque)push(a int){
	*d = append(*d,a)
}

func (d *deque)pop(){
	*d = (*d)[:len(*d)-1]
}

func (d *deque)peek()int{
	return (*d)[len(*d)-1]
}

func (d *deque)lpeek()int{
	return (*d)[0]
}

func (d *deque)lpop(){
	*d = (*d)[1:len(*d)]
}

func pushaelement(d *deque,a int,out []int,curidx int)([]int){
	cnt := 0
	for len(*d)>0 {
		top := d.peek()
		if top < a{
			d.pop()
			out = out[:len(out)-1]
			cnt++
		}else{
			break
		}
	}
	d.push(a)
	newout := append(out,curidx)
	return newout
}

func popelement(d *deque,out []int,curidx int)([]int){
	if curidx == out[0]{
		d.lpop()
		out = out[1:]
	}
	return out
}

func maxSlidingWindow(nums []int, k int) []int {
	d := new(deque)
	out := []int{}
    for i:=0;i<k;i++{
		out = pushaelement(d,nums[i],out,i)
	}
	result := []int{}
	for i:=k;i<len(nums);i++{
		result = append(result,d.lpeek())
		out = popelement(d,out,i-k)
		out = pushaelement(d,nums[i],out,i)
	}
	return result
}


func main() {
	fmt.Println("hello world")
}
