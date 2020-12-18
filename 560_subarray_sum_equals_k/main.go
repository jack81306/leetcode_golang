package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	table := make(map[int]int)
	table[0] = 1
	presum := 0
	cnt := 0
	for i:=0;i<len(nums);i++{
		presum += nums[i]
		target := presum-k
		v,ok := table[target]
		if ok {
			cnt += v
		}
		v,ok = table[presum]
		if ok{
			table[presum] = v+1
		}else{
			table[presum] = 1
		}
	}
	return cnt
}

func main() {
	fmt.Println("hello world")
}