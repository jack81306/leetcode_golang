package main

import (
	"fmt"
)



func findUnsortedSubarray(nums []int) int {
	if len(nums)==1{
		return 0
	}
	big,small := nums[0],nums[len(nums)-1]
	start,end := -1,-2
	for i:=1;i<len(nums);i++{
		big = max(big,nums[i])
		small = min(small,nums[len(nums)-i-1])
		if big > nums[i]{
			end = i
		}
		if small < nums[len(nums)-i-1]{
			start = len(nums)-i-1
		}
	}
	return end-start+1
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}

func main() {
	fmt.Println("hello world")
}
