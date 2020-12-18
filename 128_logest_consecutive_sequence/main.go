package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	result := make(map[int]bool)
	for _,v := range nums {
		result[v] = true
	}
	max := 1
	for _,v := range nums{
		_,ok := result[v-1]
		if !ok {
			cnt := 1
			cur := v+1
			for {
				_,ok := result[cur]
				if ok{
					cur++
					cnt++
				}else{
					break
				}
			}
			if cnt>max{
				max = cnt
			}
		}
	}
	return max
}

func main() {
	fmt.Println("hello world")
}
