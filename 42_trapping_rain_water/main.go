package main

import (
	"fmt"
)

//檢查每個節點，找出每個節點左邊最大值以及右邊最大值
//若 min(leftmax,rightmax) 比該節點大，那此節點至少可裝min(leftmax,rightmax)減去該節點值
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	leftmax := make([]int,len(height))
	rightmax := make([]int,len(height))
	leftmax[0] = height[0]
	rightmax[len(rightmax)-1] = height[len(rightmax)-1]
	for i:=1;i<len(height);i++{
		if height[i] > leftmax[i-1] {
			leftmax[i] = height[i]
		}else{
			leftmax[i] = leftmax[i-1]
		}
	}
	for i:=len(height)-2;i>=0;i--{
		if height[i] > rightmax[i+1] {
			rightmax[i] = height[i]
		}else{
			rightmax[i] = rightmax[i+1]
		}
	}
	result := 0
	for i:=1;i<len(height)-1;i++{
		 result += cantrap(leftmax[i-1],rightmax[i+1],height[i]) 
	}
	return result
}

func cantrap(l,r,mid int)int{
	if mid >= l || mid >= r {
		return 0
	}else{
		if r>l{
			return l-mid
		}else{
			return r-mid
		}
	}
}

func main()  {
	fmt.Println("hello world")
}