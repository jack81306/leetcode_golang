package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals,func(i,j int)bool{
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	for len(intervals)>1{
		if intervals[1][0] <= intervals[0][1]{
			tmp := [][]int{{intervals[0][0],max(intervals[0][1],intervals[1][1])}}
			intervals = append(tmp,intervals[2:]...)
		}else{
			result = append(result,intervals[0])
			intervals = intervals[1:]
		}
	}
	result = append(result,intervals...)
	return result
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}

func main() {
	fmt.Println("hello world")
}
