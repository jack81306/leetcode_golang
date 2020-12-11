package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//recursive
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
		return nil
	}
	root :=  preorder[0]
	left:=0
	for left<len(inorder)&&inorder[left]!=root{
		left++
	}
	r := new(TreeNode)
	r.Val = root
	r.Left = buildTree(preorder[1:1+left],inorder[:left])
	r.Right = buildTree(preorder[1+left:],inorder[1+left:])
	return r
}


func main()  {
	fmt.Println("hello world")
}