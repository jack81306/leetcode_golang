package main

import (
	"fmt"
)

func findLength(A []int, B []int) int {
	table := make([][]int,len(A)+1)
	for i,_ := range table{
		table[i] = make([]int,len(B)+1)
	}
	
	m := 0
	for i:=1;i<=len(A);i++{
		for j:=1;j<=len(B);j++{
			if A[i-1]==B[j-1]{
				table[i][j] = table[i-1][j-1]+1
				if table[i][j]>m{
					m = table[i][j]
				}
			}
		}
	}
	return m
}



func main() {
	fmt.Println("hello world")
}