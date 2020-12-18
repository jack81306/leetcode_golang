package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
    if len(matrix)==0{
        return 0
    }
	his := make([]int,len(matrix[0]))
	for i:=0;i<len(matrix[0]);i++{
		if matrix[0][i]=='1'{
			his[i]++
		}
	}
	res := largestRectangleArea(his)
	for i:=1;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]=='1'{
                his[j] += 1
            }else{
                his[j] =0
            }
        }
        res = max(res,largestRectangleArea(his))
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


func largestRectangleArea(heights []int) int {
	heights = append(heights,0)
	stack := []int{-1}
	res := 0
	for i,v := range heights{
        for len(stack)>1&&v < heights[stack[len(stack)-1]]{
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




func main() {
	fmt.Println("hello world")
}
