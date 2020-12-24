package main

import (
	"fmt"
)

func solve(board [][]byte)  {
    for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j] == 'O'{
				if dfs(board,i,j){
					change(board,i,j,'X')
				}
			}
		}
	}

	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j] == '*'{
				change(board,i,j,'O')
			}
		}
	}
}

func dfs(board [][]byte,i int,j int)bool{
	if i<0 || i>=len(board)||j<0||j>=len(board[0]){
		return false
	}
	if board[i][j] == 'X'{
		return true
	}
	if board[i][j] == '*'{
		return true
	}
	board[i][j] = '*'
	return dfs(board,i+1,j)&&dfs(board,i-1,j)&&dfs(board,i,j+1)&&dfs(board,i,j-1)
}

func change(board [][]byte,i int,j int,b byte){
	if i<0 || i>=len(board)||j<0||j>=len(board[0]){
		return 
	}
	if board[i][j] == '*'{
		board[i][j] = b
		change(board,i+1,j,b)
		change(board,i-1,j,b)
		change(board,i,j+1,b)
		change(board,i,j-1,b)
	}
	return 
}

func main() {
	fmt.Println("hello world")
}

