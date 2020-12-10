package main

import (
	"fmt"
)

//dp : 當root為某個點時，左子樹個數乘上右子樹個數
func numTrees(n int) int {
	result := make([]int, 19)
	result[0], result[1] = 1, 2
	for i := 2; i < 19; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i] += result[i-1]
				continue
			}
			result[i] += result[j-1] * result[i-j-1]
		}
	}
	return result[n-1]
}

func main() {
	fmt.Println("hello world")
}
