package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }

 func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}else{
		a := maxDepth(root.Left)
		b := maxDepth(root.Right)
		if a>b {
			return a+1
		}else {
			return b+1
		}
	}
}

func main()  {
	fmt.Print("hello world")
}