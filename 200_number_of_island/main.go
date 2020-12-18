package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	cnt := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j] == '1' {
				getisland(grid,i,j)
				cnt++
			}
		}
	}
	return cnt
}

func getisland(grid [][]byte,i int,j int){
	grid[i][j] = '0'
	lenx,leny := len(grid[0]),len(grid)
	if j+1 < lenx && grid[i][j+1] == '1' {
		getisland(grid,i,j+1)
	}
	if i+1<leny && grid[i+1][j] == '1'{
		getisland(grid,i+1,j)
	}
	if j-1 >= 0  && grid[i][j-1] == '1' {
		getisland(grid,i,j-1)
	}
	if i-1 >=0 && grid[i-1][j] == '1'{
		getisland(grid,i-1,j)
	}

}


func main() {
	fmt.Println("hello world")
}