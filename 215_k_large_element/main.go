package main

import (
	"fmt"
	"sort"
)
//排序
//另解:快速排序法
func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)
    return nums[len(nums)-k]
}

func main()  {
	fmt.Println("hello world")
}