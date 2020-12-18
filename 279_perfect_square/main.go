package main

import (
	"fmt"
)

func numSquares(n int) int {
	table := make([]int,n+1)
	table[0] = 0
	for i:=1;i<len(table);i++{
		min := -1
		for j:=1;i-(j*j)>=0;j++{
			min = mymin(min,table[i-j*j])
		}
		table[i] = min+1
	}
	return table[n]
}

func mymin(a,b int)int{
	if a == -1 {
		return b
	}else{
		if a > b {
			return b
		}else{
			return a
		}
	}
}


func main() {
	fmt.Println("hello world")
}