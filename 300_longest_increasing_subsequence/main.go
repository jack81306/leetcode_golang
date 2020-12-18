package main

import (
	"fmt"
)


func lengthOfLIS(nums []int) int {
	result := make([]int,len(nums))
	result[0] = nums[0]
	start,end := 0,0
	for i:=1;i<len(nums);i++{
		res := insert(result,start,end,nums[i])
		if res==-1{
			result[end+1] = nums[i]
			end++
		}else{
			result[res] = nums[i]
		}
	}
	return end+1
}


func insert(res []int,start int,end int,target int)int{
	if start > end{
		return -1
	}
	mid := int((start+end)/2)
	if start == end {
		if target <= res[start]{
			return start
		}
		return -1
	}
	if res[mid] == target{
		return mid
	}else if res[mid] < target{
		return insert(res,mid+1,end,target)
	}else{
		leftmax := insert(res,start,mid-1,target)
		if leftmax==-1{
			return mid
		}else{
			return leftmax
		}
	}
}

func main() {
	t := []int{4,10,4,3,8,9}
	fmt.Println(lengthOfLIS(t))
}