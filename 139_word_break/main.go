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
    tmp:=  dp(s,wordDict,0)
	return tmp
}

func dp(s string,wordDict []string,start int) []string{
	if isdo[start]{
		return table[start]
	}

	res := []string{}
	for	_,v := range wordDict{
		if isprefix(s[start:],v){
            if len(s[start:])==len(v){
				res = append(res,v)
			}else{
                if start+len(v)>=len(s){
                    continue
                }
                tmp := dp(s,wordDict,start+len(v))
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
	s := "hello world"
	s2 := "heelo"
	fmt.Println(isprefix(s,s2))
}
