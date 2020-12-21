package main

import (
	"fmt"
)
func findRedundantConnection(edges [][]int) []int {
	table := make([]int,len(edges)+1)
	for i,_ := range table{
		table[i] = i
	}
	result := []int{}
	for i:=0;i<len(edges);i++{
		first,second := edges[i][0],edges[i][1]
		for table[first]!=first{
			first = table[first]
		}
		for table[second]!=second{
			second = table[second]
		}
		if first == second{
			result = edges[i]
		}else{
			tmp := edges[i][1]
			for tmp!=second{
				tmp2 := table[tmp]
				table[tmp] = first
				tmp = tmp2
			}
			table[tmp] = first
		}
	}
	return result
}

func main() {
	fmt.Println("hello world")
}
