package main

import (
	"fmt"
)

func sortColors(nums []int) {
	table := make([]int, 3)
	for _, v := range nums {
		table[v]++
	}
	idx := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < table[i]; j++ {
			nums[idx] = i
			idx++
		}
	}
}

func main() {
	fmt.Println("hello world")
}
