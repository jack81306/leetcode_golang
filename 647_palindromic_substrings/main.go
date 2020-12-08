package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	f,f2 := getsubstring(s)
	for j:=0;j<len(s);j++{
		for i:=0;i<len(s)-j;i++{
			f(i,i+j)
		}
	}
	return f2()
}

//dp  使用string[start+1][end-1]來檢查string[start][end]是否回文
//go closure練習
func getsubstring(s string) (func(int,int),func() int) {
	table := make([][]int,len(s))
	for i := range table {
		table[i] = make([]int,len(s))
	}

	for i:=0;i<len(s);i++{
		table[i][i] = 1
	}

	f := func(start int,end int){
		if start >= end {
			return
		}
		if s[start] == s[end] {
			if end - start == 1{
				table[start][end] = 1
			}else{
				if  table[start+1][end-1] == 1 {
					table[start][end] = 1
				}else{
					table[start][end] = 0
				}
			}
		}
	}

	f2 := func() int{
		result := 0
		for i:=0;i<len(s);i++{
			for j:=0;j<len(s);j++{
				result += table[i][j]
			}
		}
		return result
	}
	return f,f2
}

func main()  {
	fmt.Println(countSubstrings("abc"))
}