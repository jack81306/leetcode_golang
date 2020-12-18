package main

import (
	"fmt"
)

func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	notro,ro := make([]int,len(nums)),make([]int,len(nums))
	notro[0] = 0
	ro[0] = nums[0]
	for i:=1;i<len(nums);i++{
		ro[i] = notro[i-1] + nums[i]
		notro[i] = max(notro[i-1],ro[i-1])
	}
	return max(ro[len(ro)-1],notro[len(notro)-1])
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}

func main() {
	fmt.Println("hello world")
}