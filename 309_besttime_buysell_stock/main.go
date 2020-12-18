package main

import (
	"fmt"
)

type buyprofit struct{
	buyprice int

}

func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	s0 := make([]int,len(prices))
	s1 := make([]int,len(prices))
	s2 := make([]int,len(prices))

	s0[0] = 0
	s1[0] = -prices[0]
	var tmp int32 = 1
	s2[0] = int(tmp<<31)
	for i:=1;i<len(prices);i++{
		s0[i] = max(s0[i-1],s2[i-1])
		s1[i] = max(s1[i-1],s0[i-1]-prices[i])
		s2[i] = s1[i-1]+prices[i]
	}
	return max(s0[len(prices)-1],s2[len(prices)-1],s2[len(prices)-1])
}

func max(a ...int)int{
	max := a[0]
	for i:=1;i<len(a);i++{
		if a[i] > max {
			max = a[i]
		}
	}
	return max 
}

func main() {
	fmt.Println("hello world")
}