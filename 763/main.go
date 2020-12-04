package main

import (
	"fmt"
)

func partitionLabels(S string) []int {
	type pair struct{
		start int
		end int
	}
	all := []pair{}
	result := []int{}
	for i:=0 ; i < 26;i++ {
		all = append(all,pair{start:-1,end:0})
	} 

	for i:=0;i<len(S);i++{
		c := S[i] - 'a'
		if all[c].start == -1 {
			all[c].start = i
		}
		all[c].end = i
	}

	cnt := 0
	for i:=1;i<len(S);i++{
		cnt += 1 
		iscut := true
		for j:=0;j<26;j++{
			if all[j].start == -1 {
				continue
			}
			if all[j].start < i && all[j].end >= i {
				iscut = false
				break
			}
		}
		if iscut {
			result = append(result,cnt)
			cnt = 0
		}
	}
	result = append(result,cnt+1)
	return result
}

func  main()  {
	partitionLabels("ababcbacadefegdehijhklij")
}