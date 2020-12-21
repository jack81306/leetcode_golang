package main

import (
	"fmt"
)
//You can assume that you can always reach the last index.
//所以才不用檢查是否到最後
func jump(nums []int) int {
	reach,i,mincnt,curfar := 0,0,0,0
	n := len(nums)
	for ;i<n&&i<=reach;i++{
		if nums[i]+i > reach{
			reach = nums[i]+i
		}
		if curfar == i{
			mincnt++
			curfar = reach
		}
	}

	return mincnt
}

func min(a ...int)int  {
	m := a[0]
	for _,v := range a{
		if v<m{
			m = v
		}
	}
	return m
}


func main() {
	fmt.Println("hello world")
}
