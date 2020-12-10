package main

import (
	"fmt"
)

//dp
//dp[i,j] = dp[i-1,j] + dp[i,j-1]
//可用一維振烈節省空間
func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		result[i][0] = 1
	}
	for i := 0; i < n; i++ {
		result[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	return result[m-1][n-1]
}

func main() {
	fmt.Println("hello world")
}
