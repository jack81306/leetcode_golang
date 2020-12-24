package main

import (
	"fmt"
)

//reference:https://blog.csdn.net/fuxuemingzhu/article/details/79473171
func findMaximumXOR(nums []int) int {
	var max,mask int = 0,0
	for i:=31;i>=0;i--{
		table := make(map[int]bool)
        mask = (1<<i)|mask
		tempmax := max|1<<i
		for _,v:=range nums{
			table[mask&v]=true
		}
		for k,_ := range table{
			_,ok := table[tempmax^k]
			if ok{
				max = tempmax
				break
			}
		}
	}
	return max
}



func main() {
	fmt.Println("hello world")
}

