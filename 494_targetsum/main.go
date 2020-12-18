package main

import (
	"fmt"
	"sort"
)

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _,v := range nums{
		sum += v
	}
	sort.Ints(nums)
	positive := int((sum-S)/2+S) //may some troble
	table := make([][]int,len(nums))
	for i,_ := range table{
		table[i] = make([]int,positive+1)
	}
	
	if nums[0] == 0{
		table[0][0] = 2
	}else{
		table[0][0] = 1
		for i:=1;i<=positive;i++{
			if nums[0] == i{
				table[0][i] = 1
			}
		}
	}

	for i:=1;i<len(table);i++{
		if nums[i]==0{
			table[i][0] = table[i-1][0]*2
		}else{
			table[i][0] = table[i-1][0]
		}
	}

	for i:=1;i<len(table);i++{
		for j:=1;j<=positive;j++{
			//要拿i項或不拿
			nottake := table[i-1][j]
			take := 0
			if positive - nums[i] >= 0{
				take = table[i][positive - nums[i]]
			}		
			table[i][j] = take+nottake
		}
	}
	return table[len(nums)-1][positive]
}


func main() {
	fmt.Println("hello world")
}
