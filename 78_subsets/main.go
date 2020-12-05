package main

import (
	"fmt"
)
//透過使用二進制的01代表第幾位的元素是否出現
func subsets(nums []int) [][]int {
	total := 1
	total = total<<len(nums)
	result := [][]int{}
	for i:=0;i<total;i++{
		tmp := i
		r := []int{}
		for j:=0;j<len(nums);j++{
			if tmp&1 == 1{
				r = append(r,nums[j])
			}
			tmp = tmp>>1
		}
		result = append(result,r)
	}
	return result
}

func main()  {
	fmt.Println("hello world")
}