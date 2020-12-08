package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	f,f2 := getsum(candidates,target)
	for i:=1;i<=target;i++{
		f(i)
	}
	return f2(target)
}
//closure 練習
//dp
func getsum(candidates []int,target int)(func(int),func(int)[][]int){
	table := make([][][]int,target+1)
	table[0] = [][]int{{},}
	f := func(target int){
		r := [][]int{}
		for _,v := range candidates {
			if target - v < 0 || len(table[target-v]) ==0 {
				continue
			}
			for _,v2:=range table[target-v]{
				//透過限定由小到大來避免重複
				if len(v2) > 0 {
					if v<v2[len(v2)-1]{
						continue
					}	
				}
				//以免共用相同的記憶體空間
				tmp := []int{}		
				tmp = append(tmp,v2...)
				tmp = append(tmp,v)
				r = append(r,tmp)
			}
		}
		table[target] = r
	}
	f2 := func(target int) [][]int {
		return table[target]
	}
	return f,f2
}

func main()  {
	fmt.Println("hello world")
}