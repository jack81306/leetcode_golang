package main

import (
	"fmt"
)


//視為從左邊找到第一個較小與右邊第一個較小
//使用stack紀錄
func largestRectangleArea(heights []int) int {
	heights = append(heights,0)
	stack := []int{-1}
	res := 0
	for i,v := range heights{
        for v < heights[stack[len(stack)-1]]{
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[val]
			w := i-stack[len(stack)-1]-1
			res = max(res,h*w)
		}
		stack = append(stack,i)
	}
	return res
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
