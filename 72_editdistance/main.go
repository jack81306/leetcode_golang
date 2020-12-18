package main

import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	if len(word1)==0 || len(word2)==0{
		return len(word1)+len(word2)
	}
	table := make([][]int,len(word1)+1)
	for i:=0;i<len(table);i++{
		table[i] = make([]int,len(word2)+1)
	}
	table[0][0] = 0
	for i:=1;i<len(table[0]);i++{
		table[0][i] = table[0][i-1]+1
	}

	for i:=1;i<len(table);i++{
		table[i][0] = table[i-1][0]+1
	}
	for i:=1;i<len(table);i++{
		for j:=1;j<len(table[0]);j++{
			if word1[i-1] == word2[j-1]{
				table[i][j] = min(table[i-1][j]+1,table[i][j-1]+1,table[i-1][j-1])
			}else{
				table[i][j] = min(table[i-1][j]+1,table[i][j-1]+1,table[i-1][j-1]+1)
			}
		}
	}
	return table[len(word1)][len(word2)]
}

func min(all ...int)int{
	m := all[0]
	for _,v:=range all{
		if v < m{
			m =v
		}
	}
	return m
}

func main() {
	fmt.Println("hello world")
}
