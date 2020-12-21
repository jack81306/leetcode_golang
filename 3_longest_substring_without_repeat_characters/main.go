package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	now := make(map[byte]int)
	l,max := -1,0
	for i:=0;i<len(s);i++{
		v,ok := now[s[i]]
        if ok && v>l {
            l = now[s[i]]
        }
        now[s[i]] = i
		if i-l > max{
			max = i-l
		}
	}
	return max
}
func check(m map[byte]int)bool{
	for _,v := range m{
		if v > 1{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("hello world")
}
