package main

import (
	"fmt"
)

var table [][]string
var isdo []bool


func wordBreak(s string, wordDict []string) []string {
	table = make([][]string,len(s))
	isdo = make([]bool,len(s))
	for i,_ := range table{
		table[i] = []string{}
	}
	return dp(s,wordDict,0)
}

func dp(s string,wordDict []string,start int) []string{
	if isdo[start]{
		return table[start]
	}

	res := []string{}
	for	_,v := range wordDict{
		if isprefix(s[start:],v){
			if len(s)==len(v){
				res = append(res,v)
			}else{
				tmp := wordBreak(s[len(v):],wordDict)
				if len(tmp)>0{
					for _,val := range tmp{
						res = append(res,v+" "+val)
					}
				}
			}
		}
	}
	table[start] = res
	isdo[start] = true
	return res
}
func isprefix(s,s2 string)bool{
	if len(s2) > len(s){
		return false
	}
	if s[:len(s2)] == s2{
		return true
	}
	return false
}

func main() {
	fmt.Println("hello world")
}
