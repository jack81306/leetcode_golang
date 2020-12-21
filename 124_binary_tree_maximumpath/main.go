package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func maxPathSum(root *TreeNode) int {
    if root==nil{
		return 0
	}
	res,_ := maxpath(root)
	return res
}

func maxpath(root *TreeNode)(int,int){
	if root==nil{
		return -1000000000,0
	}
	l,lh := maxpath(root.Left)
	r,rh := maxpath(root.Right)
    return max(l,r,root.Val,root.Val+lh,root.Val+rh,root.Val+lh+rh),max(lh,rh,0)+root.Val
}

func max(a ...int)int {
    m := a[0]
    for _,v := range a{
        if v > m{
            m = v
        }
    }
    return m
}

func main() {
	fmt.Println("hello world")
}
