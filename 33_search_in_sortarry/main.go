package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	p := findpivot(nums,0,len(nums)-1)
	if nums[p] == target{
		return p
	}
	if nums[len(nums)-1] >= target{
		return bs(nums,target,p+1,len(nums)-1)
	}else{
		return bs(nums,target,0,p-1)
	}
}

func bs(nums[]int,target int, start int,end int)int{
	if start == end {
		if target == nums[start]{
			return start
		}else{
			return -1
		}
	}
	if start > end{
		return -1
	}
	mid := int((start+end)/2)
	if target == nums[mid]{
		return mid
	}else if target < nums[mid]{
		return bs(nums,target,start,mid-1)
	}else{
		return bs(nums,target,mid+1,end)
	}
}

func findpivot(nums []int,start int,end int)int{
	if start==end{
		return start
	}
	mid := int((start+end)/2)
	if nums[end] < nums[mid]{
		// pivot at right
		return findpivot(nums,mid+1,end)
	}else if nums[start] > nums[mid]{
		//pivot at left
		return findpivot(nums,start,mid)
	}
	return start
}
func main() {
	s := []int{4,5,6,7,0,1,2}
	fmt.Println(findpivot(s,0,6))
}
