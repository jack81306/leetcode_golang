package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func maxSubArray(nums []int) int {
	if len(nums)==0{
		return 0
	}
	result := make([]int,len(nums))
	result[0] = nums[0]
	max := result[0]
	for i:=1;i<len(nums);i++{
		if result[i-1] < 0{
			result[i] = nums[i]
		}else{
			result[i] = result[i-1]+nums[i]
		}
		if result[i] > max{
			max = result[i]
		}
	}
	return max
}