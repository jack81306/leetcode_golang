package main

import (
	"fmt"
)

//可用交換 代替複製 節省記憶體空間
// no recursion:從0個元素開始 依序插入第n個元素到第n-1已完成的排列中
func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{{nums[0]},}
    }
	result := [][]int{}
	c := make([]int,len(nums))
	copy(c,nums)
    for i := 0;i<len(c);i++{
		topermute := []int{}
		topermute = append(topermute,c[:i]...)
		topermute = append(topermute,c[i+1:]...)
		r := permute(topermute)
        for _,s := range r {
            result = append(result,append([]int{nums[i]},s...))
		}
	}	
	return result
}
func main()  {
	n := []int{1,2,3}
	fmt.Println(permute(n))
}