package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//easy dp dp[i,j] = min(dp[i-1,j],dp[i,j-1])+grid[i,j]
func minPathSum(grid [][]int) int {
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	fmt.Println("hello world")
}
