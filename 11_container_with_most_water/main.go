package main

import (
	"fmt"
)

func maxArea(height []int) int {
	water := 0
	i, j := 0, len(height)-1
	for i < j {
		h := min(height[a], height[b])
		water = (j - i) * h
		for height[i] <= h && i < j {
			i++
		}
		for height[j] <= h && i < j {
			j--
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println("hello world")
}
