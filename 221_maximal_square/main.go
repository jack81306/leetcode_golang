package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
    if len(matrix)==0{
		return 0
	}
	table := make([][]int,len(matrix))
	for i,_ := range table{
		table[i] = make([]int,len(matrix[0]))
	}
	for i:=0;i<len(table[0]);i++{
		if matrix[0][i] == '1'{
			table[0][i] = 1
		}
	}
	for i:=0;i<len(table);i++{
		if matrix[i][0] == '1'{
			table[i][0] = 1
		}
	}
	for i:=1;i<len(table);i++{
		for j:=1;j<len(table[0]);j++{
			if matrix[i][j] == '1'{
				table[i][j] = min(table[i-1][j],table[i-1][j-1],table[i][j-1])+1
			}
		}
	}
	res := 0
	for i:=0;i<len(table);i++{
		for j:=0;j<len(table[0]);j++{
			if table[i][j]>res{
				res = table[i][j]
			}
		}
	}
	return res
}

func min(a ...int)int{
	m := a[0]
	for _,v := range a{
		if v<m{
			m = v
		}
	}
	return m
}

func main() {
	fmt.Println("hello world")
}
