package main

import (
	"fmt"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//easy recursion 先佐在右
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)
	result := []int{}
	if l!= nil {
		result = append(result,l...)
	}
	result = append(result,root.Val)
	if r!= nil {
		result = append(result,r...)
	}
	return result
}

func main()  {
	fmt.Println("hello world")
}