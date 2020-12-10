package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	result := make([]int, len(prices))
	maxprofit := 0
	result[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < result[i-1] {
			result[i] = prices[i]
		} else {
			result[i] = prices[i-1]
			if prices[i]-result[i] > maxprofit {
				maxprofit = prices[i] - result[i]
			}
		}
	}
	return maxprofit
}

func main() {
	fmt.Println("hello world")
}
