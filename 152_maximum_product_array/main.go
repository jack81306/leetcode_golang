package main

import (
	"fmt"
)


func maxProduct(nums []int) int {
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	positive,negative := make([]int,len(nums)),make([]int,len(nums))
	if nums[0] > 0{
		positive[0] = nums[0]
	}else{
		negative[0] = nums[0]
	}
	for i:=1;i<len(nums);i++{
		if nums[i]>0{
			positive[i] = max(positive[i-1]*nums[i],nums[i])
			negative[i] = negative[i-1]*nums[i]
		}else if nums[i]<0{
			positive[i] = negative[i-1]*nums[i]
			negative[i] = min(positive[i-1]*nums[i],nums[i])
		}
	}
	max := positive[0]
	for i:=0;i<len(positive);i++{
		if positive[i]>max{
			max = positive[i]
		}
	}
	return max
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
