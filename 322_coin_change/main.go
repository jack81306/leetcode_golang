package main

import (
	"fmt"
)
func coinChange(coins []int, amount int) int {
	table := make([]int,amount+1)
	table[0] = 0
	var tmp int32 = 1
	intmax := int(tmp<<31-1)
	for i:=1;i<=amount;i++{
		table[i] = intmax
		for _,v := range coins{
			if i-v >= 0{
				table[i] = min(table[i],table[i-v]+1)
			}
		}
	}
	if table[amount]==intmax{
		return -1
	}
	return table[amount]
}

func min(a,b int)int{
	if a<b{
		return a
	}else{
		return b
	}
}

func main() {
	fmt.Println("hello world")
}
