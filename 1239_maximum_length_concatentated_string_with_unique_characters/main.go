package main

import (
	"fmt"
)

func maxLength(arr []string) int {
	input := []int{}
	for _,v := range arr{
		val := 0
        unique := true
		for i:=0;i<len(v);i++{
            target := 1<<(v[i]-'a')
            if val & target != 0{
                unique = false
                break
            }else{
                val = val|target
            }
		}
        if unique{
            input = append(input,val)   
        }
	}
	res := dfs(input)
	max := 0
	for _,v := range res{
		cnt := 0
		for v>0{
			cnt++
			v=v&(v-1)
		}
		if cnt>max{
			max=cnt
		}
	}
	return max
}

func dfs(arr []int)[]int{
    if len(arr)==0{
        return []int{}
    }
	if len(arr)==1{
		return []int{arr[0]}
	}

	res := []int{arr[0]}
	tmp := dfs(arr[1:])
	for _,v := range tmp{
		val := arr[0]&v
		if val == 0{
			res = append(res,arr[0]|v)
        }
        res = append(res,v)
	}
	return res
}

func main() {
	fmt.Println("hello world")
}

