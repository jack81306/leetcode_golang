package main

import (
	"fmt"
)

func countBits(num int) []int {
	result := []int{0}
	key := 1 
	count := 0
	for i := 0;i<num;i++{
		if key == count {
			count = 0
			key = key*2
		}
		result = append(result,1+result[count])
		count++
	}
	return result
}

func main()  {
	fmt.Println(countBits(5))
}