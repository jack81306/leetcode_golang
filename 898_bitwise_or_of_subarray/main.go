package main

import (
	"fmt"
)

func subarrayBitwiseORs(arr []int) int {
	if len(arr)==0{
		return 0
	}
	res,prev,now := make(map[int]bool),make(map[int]bool),make(map[int]bool)
	res[arr[0]] = true
	prev[arr[0]] = true
	for i:=1;i<len(arr);i++{
		now = make(map[int]bool)
		now[arr[i]] = true
		for key,_ := range prev{
			now[key|arr[i]]=true
		}
		for key,_ := range now{
			res[key] = true
		}
		prev = now
	}
	return len(res)
}

func main() {
	fmt.Println("hello world")
}

