package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	fromhead := make([]int,len(nums))
	fromtail := make([]int,len(nums))
	a,b := 1,1
	l := len(nums)
	for i:=0;i<l;i++{
		a *= nums[i]
		b *= nums[l-1-i]
		fromhead[i] = a
		fromtail[l-1-i] = b
	}
	result := make([]int,len(nums))
	result[0] = fromtail[1]
	result[l-1] = fromhead[l-2]
	for i:=1;i<l-1;i++{
		result[i] = fromhead[i-1]*fromtail[i+1]
	}
	return result
}

func main()  {
	fmt.Println("hello world")
}