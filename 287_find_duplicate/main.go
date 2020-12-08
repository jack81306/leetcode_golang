package main

import (
	"fmt"
	"sort"
)



func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0;i<len(nums)-1;i++{
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main()  {
	fmt.Println("hello world")
}