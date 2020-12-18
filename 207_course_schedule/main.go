package main

import (
	"fmt"
)


func canFinish(numCourses int, prerequisites [][]int) bool {
	table := make([][]bool,numCourses)
	ispick := make([]bool,numCourses)
	for i,_ := range table{
		table[i] = make([]bool,numCourses)
	}

	for _,v:= range prerequisites{
		table[v[0]][v[1]] = true
	}

	for checkallpick(ispick){
		find := false
		for i:=0;i<numCourses;i++{
			if !checkrequest(table[i]){
				find = true
				for j:=0;j<numCourses;j++{
					table[j][i] = false
				}
				ispick[i] = true
				break
			}
		}
		if !find{
			return false
		}
	}
	return true
}

func checkrequest(a []bool) bool {
	for _,v := range a{
		if v == true{
			return true
		}
	}
	return false
}

func checkallpick(a []bool) bool {
	for _,v := range a{
		if v == false{
			return true
		}
	}
	return false
}


func main() {
	fmt.Println("hello world")
}