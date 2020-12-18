package main

import (
	"fmt"
)



func exist(board [][]byte, word string) bool {
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if dfs(board,word,i,j,0){
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte,word string,x int,y int,cnt int)bool{
	if x<0 || x>=len(board)||y<0||y>=len(board[0]){
		return false
	}

	if board[x][y]!=word[cnt]{
		return false
	}
	if cnt==len(word)-1{
		return true
	}
	tmp := board[x][y]
	board[x][y] = '%'
	res := dfs(board,word,x-1,y,cnt+1)||dfs(board,word,x+1,y,cnt+1)||dfs(board,word,x,y+1,cnt+1)||dfs(board,word,x,y-1,cnt+1)
	board[x][y]=tmp
	return res
}


func main() {
	fmt.Println("hello world")
}
