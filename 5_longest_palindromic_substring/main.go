package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	table := make([][]bool,len(s))
	for i,_ := range table{
		table[i] = make([]bool,len(s))
	}
	for i:=0;i<len(table);i++{
		table[i][i]=true
	}
	maxlen := 1
	maxstr := string(s[0])
	for j:=1;j<len(table);j++{
		for i:=0;i+j<len(table);i++{
			if s[i]==s[j]{
				if j-i == 1{
					table[i][j] = true
				}else{
					table[i][j] = table[i+1][j-1]
				}
			}else{
				table[i][j] = false
			}
			if table[i][j] == true && j-i+1 > maxlen{
				maxlen = j-i+1
				maxstr = s[i:j+1]
			}
		}
	}
	return maxstr
}

func main() {
	fmt.Println("hello world")
}
