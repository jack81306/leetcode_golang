package main

import (
	"fmt"
)

func findCircleNum(M [][]int) int {
	table := make([]int,len(M))
	for i,_ := range table{
		table[i]=i
	}
	for i:=0;i<len(M);i++{
		for j:=i;j<len(M);j++{
			if M[i][j]==1{
				personi,personj,groupi,groupj := i,j,table[i],table[j]
				for personi!=groupi{
					personi = groupi
					groupi = table[personi]
				}
				for personj!=groupj{
					personj = groupj
					groupj = table[personj]
				}
				if groupi == groupj{
					continue
				}else{
					tmpj := j
					for tmpj != personj{
						tmp := table[tmpj]
						table[tmpj] = groupi
						tmpj=tmp
					}
					table[personj] = groupi
				}
			}
		}
	}
	cnt := 0
	for i,v := range table{
		if i==v{
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println("hello world")
}

