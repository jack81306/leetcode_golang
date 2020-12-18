package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for i,v := range nums{
		val,ok := result[target-v]
		if ok{
			return []int{val,i}
		}else{
			result[v] = i
		}
	} 
	return []int{}
}


func main() {
	fmt.Println("hello world")
}
