package main

import (
	"fmt"
)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q{
		return root
	}
	l := lowestCommonAncestor(root.Left,p,q)
	r := lowestCommonAncestor(root.Right,p,q)
	if l==nil && r == nil{
		return nil
	}else if l==nil{
		return r
	}else if r==nil{
		return l
	}else{
		return root
	}
}


func main() {
	fmt.Println("hello world")
}