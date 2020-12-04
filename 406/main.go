package main

import (
	"sort"
)

// func swap(indexi int,indexj int,a [][]int)[][]int{
// 	x := a[indexi][0]
// 	y := a[indexi][1]
// 	a[indexi][0] = a[indexj][0]
// 	a[indexi][1] = a[indexj][1]
// 	a[indexj][0] = x
// 	a[indexj][1] = y
// 	return a 
// }

func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people,func(i,j int) bool{
		if  people[i][0] == people[j][0]{
			return  people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	result := [][]int{}
	result = append(result,people[0])
	for i:=1 ; i<len(people);i++{
		if people[i][1] == len(result){
			result = append(result,people[i])
			continue
		}
		result = append(result[:people[i][1]+1],result[people[i][1]:]...)
		result[people[i][1]] = people[i]
	}
	return result
}

func main()  {
	test := [][]int{
		{7,0},
		{4,4},
		{7,1},
		{5,0},
		{6,1},
		{5,2},
	}
	reconstructQueue(test)
}