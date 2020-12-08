package main

import (
	"fmt"
)
//限定inplace algo
//對角翻一次 在reverse一次
func rotate(matrix [][]int)  {
	l := len(matrix[0])
	for i:=0;i<l;i++{
		for j:=i;j<l;j++{
			swap(matrix,i,j,j,i)
		}
	}
	for i:=0;i<l;i++{
		for j:=0;j<int(l/2);j++{
			swap(matrix,i,j,i,l-j-1)
		}
	}
}

func swap(matrix [][]int,dx int,dy int,dx2 int,dy2 int){
	tmp := matrix[dx][dy]
	matrix[dx][dy] = matrix[dx2][dy2]
	matrix[dx2][dy2] = tmp
}

func main()  {
	fmt.Println("hello world")
}