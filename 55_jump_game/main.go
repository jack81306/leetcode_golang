package main

import (
	"fmt"
)

var table []int

func canJump(nums []int) bool {
	reach := 0
	for i:=0;i<len(nums)&&i<=reach;i++{
		if i+nums[i] > reach {
			reach = i+nums[i]
		}
	}
	return reach == len(nums)-1
}

func main() {
	fmt.Println("hello world")
}
