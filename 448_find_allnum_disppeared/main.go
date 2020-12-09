package main

import (
	"fmt"
)

func swap(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func findDisappearedNumbers(nums []int) []int {
	for idx := range nums {
		for nums[idx] != idx+1 {
			swap(nums, idx, nums[idx]-1)
			if nums[idx] == nums[nums[idx]-1] {
				break
			}
		}
	}
	result := []int{}
	for idx := range nums {
		if nums[idx] != idx+1 {
			result = append(result, idx+1)
		}
	}
	return result
}

func main() {
	fmt.Println("hello world")
}
