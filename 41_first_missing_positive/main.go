package main

import (
	"fmt"
)
func firstMissingPositive(nums []int) int {
	l:=len(nums)
	if l==0{
		return 1
	}
	for i:=0;i<l;i++{
		for 0<=nums[i]-1 && nums[i]-1<l{
			if nums[i] == nums[nums[i]-1]{
				break
			}
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		}
	}
    for i:=0;i<l;i++{
        if nums[i]!=i+1{
            return i+1
        }
    }
    return len(nums)+1
}

func main() {
	fmt.Println("hello world")
}
