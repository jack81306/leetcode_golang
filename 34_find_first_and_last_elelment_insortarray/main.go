package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	l,r := binarysmall(nums,target,0,len(nums)-1),binarybig(nums,target,0,len(nums)-1)
	return []int{l,r}
}

func binarysmall(nums []int, target int,start int,end int)int{
	if start>end{
		return -1
	}
	if start==end{
		if nums[start]==target{
			return start
		}else{
			return -1
		}
	}
	mid := int((start+end)/2)
	if nums[mid]==target{
		l := binarysmall(nums,target,start,mid-1)
		if l!=-1{
			return l
		}else{
			return mid
		}
	}else if nums[mid]>target{
		return binarysmall(nums,target,start,mid-1)
	}else{
		return binarysmall(nums,target,mid+1,end)
	}
}

func binarybig(nums []int, target int,start int,end int)int{
	if start>end{
		return -1
	}
	if start==end{
		if nums[start]==target{
			return start
		}else{
			return -1
		}
	}
	mid := int((start+end)/2)
	if nums[mid]==target{
		l := binarybig(nums,target,mid+1,end)
		if l!=-1{
			return l
		}else{
			return mid
		}
	}else if nums[mid]>target{
		return binarybig(nums,target,start,mid-1)
	}else{
		return binarybig(nums,target,mid+1,end)
	}
}

func main() {
	fmt.Println("hello world")
}
