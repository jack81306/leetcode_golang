package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	i,j := 0,len(matrix[0])-1
	for i<len(matrix)&&j>=0{
		if matrix[i][j] == target{
			return true
		}
		if matrix[i][j] > target{
			j--
		}else{
			i++
		}
	}
	return false
}




func main() {
	fmt.Println("hello world")
}