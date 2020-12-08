package main

import (
	"fmt"
)

//Boyer-Moore Majority Vote Algorithm
func majorityElement(nums []int) int {
	count,major := 1,nums[0]
    for i:=1;i<len(nums);i++{
		if count==0{
			major = nums[i]
			count = 1
		}else{
			if nums[i]==major{
				count++
			}else{
				count--
			}
		}

	}
	return major
}

func main(){
	fmt.Println("hello world")
}