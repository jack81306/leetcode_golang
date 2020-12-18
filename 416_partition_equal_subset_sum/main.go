package main

import (
	"fmt"
)
func canPartition(nums []int) bool {
	total := 0
	for _,v := range nums{
		total += v
	}
	if total % 2 != 0 || total==0 {
		return false
	}
	table := make([][]bool,len(nums)+1)
	for i,_ := range table{
		table[i] = make([]bool,total/2)
	}
	for i:=0;i<len(table);i++{
		table[i][0] = true
	}
	for j:=1;j<len(table[0]);j++{
		table[0][j] = false
	}
	for i:=1;i<len(table);i++{
		for j:=1;j<len(table[0]);j++{
			table[i][j] = table[i-1][j]
			if j-nums[i-1] >= 0{
				table[i][j] = table[i][j] || table[i-1][j-nums[i-1]]
			}
		}
	}
	return table[len(table)-1][len(table[0])-1]
}


func main() {
	fmt.Println("hello world")
}
