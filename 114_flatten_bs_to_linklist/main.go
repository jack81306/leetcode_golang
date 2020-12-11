package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//recursive求解 
func flatten(root *TreeNode)  {
    flat(root)
}


func flat(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	l := flat(root.Left)
	r := flat(root.Right)

	root.Right = l
	root.Left = nil
	cur := root 
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = r
	return root
}



func main()  {
	fmt.Println("hello world")
}