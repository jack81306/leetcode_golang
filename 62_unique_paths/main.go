package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	var result, divisorn int64
	result = 1
	divisorn = 1
	for i := m + 1; i <= m+n-2; i++ {
		result *= i
	}
	for i := int64(1); i <= n-1; i++ {
		divisorn *= i
	}
	return result / divisorn
}

func main() {
	fmt.Println("hello world")
}
