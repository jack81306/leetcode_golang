package main

import (
	"fmt"
)

var table []int

func wordBreak(s string, wordDict []string) bool {
	table = make([]int,len(s))
	return dp(s,wordDict,0)
}

func dp(s string,wordDict []string,start int)bool{
	if table[start]!=0{
		if table[start] == 1{
			return false
		}else{
			return true
		}
	}

	l := len(s)-start
	for _,v := range wordDict{
		if len(v)<=l{
			if v==s[start:]{
				table[start] = 2
				return true
			}
			if isprefix(s[start:],v){
				if dp(s,wordDict,start+len(v)){
					table[start] = 2
					return true
				}
			}
		}
	}
	table[start] = 1
	return false
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
	s := "hello world"
	s2 := "heelo"
	fmt.Println(isprefix(s,s2))
}
