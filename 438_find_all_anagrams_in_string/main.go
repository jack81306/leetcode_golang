package main

import (
	"fmt"
)


func findAnagrams(s string, p string) []int {
    if len(s) < len(p){
		return []int{}
	}
	nowp,nows := make([]int,26),make([]int,26)
	for i:=0;i<len(p);i++{
		cs,cp := s[i],p[i]
		nows[int(cs-'a')]++
		nowp[int(cp-'a')]++
	}
	result := []int{}
	for i:=0;i<=len(s)-len(p);i++{
		if compare(nowp,nows){
			result = append(result,i)
		}
		if i+len(p) < len(s){
			nows[int(s[i]-'a')]--
			nows[int(s[i+len(p)]-'a')]++
		}
	}
	return result
}

func compare(a,b []int)bool{
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("hello world")
}
