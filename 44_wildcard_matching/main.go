package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	table := make([][]bool,len(s)+1)
	for i,_ := range table {
		table[i] = make([]bool,len(p)+1)
	}
	table[0][0] = true
    for i:=1;i<len(table[0]);i++{
        if p[i-1]=='*'{
            table[0][i] = table[0][i-1]
        }
    }
	for i:=1;i<len(table);i++{
		for j:=1;j<len(table[0]);j++{
			if s[i-1]==p[j-1]||p[j-1]=='?'{
				table[i][j] = table[i-1][j-1]
			}else{
				if p[j-1]=='*'{
					table[i][j] = table[i-1][j-1]||table[i][j-1]||table[i-1][j]
				}else{
					table[i][j]=false
				}
			}
		}
	}
    // for _,v:=range table{
    //     fmt.Println(v)
    // }
	return table[len(s)][len(p)]
}

func main() {
	fmt.Println("hello world")
}
