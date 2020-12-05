package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key int
	val int	
}

//使用map 可降在nlogn內解決
//此題可使用quick select
func topKFrequent(nums []int, k int) []int {
	all:= make(map[int]int)
	for _,val := range nums {
		v,ok := all[val]
		if ok {
			all[val] = v+1
		}else{
			all[val] = 1
		}
	}
	pairresult := []pair{}
	for key,val := range all {
		pairresult = append(pairresult,pair{key:key,val:val})
	}
	sort.Slice(pairresult,func(i,j int)bool{
		return pairresult[i].val > pairresult[j].val
	})
	result := []int{}
	for i:=0;i<k;i++{
		result = append(result,pairresult[i].key)
	}
	return result
}

func main()  {
	a := 3
	fmt.Println("hello world",a)
}