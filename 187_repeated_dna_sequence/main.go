package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	if len(s)<10{
		return []string{}
	}
	table := make(map[string]bool)
	res := []string{}
	for i:=0;i<=len(s)-10;i++{
		v,ok := table[s[i:i+10]]
		if ok{
			if v==false{
				res = append(res,s[i:i+10])
				table[s[i:i+10]]=true
			}
		}else{
			table[s[i:i+10]]=false
		}
	}
	return res
}

func main() {
	fmt.Println("hello world")
}

