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
func isSymmetric(root *TreeNode) bool {
	mirrorroot := mirror(root)
	return compare(root,mirrorroot)
}

func mirror(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	newroot := new(TreeNode)
	newroot.Val = root.Val
	newroot.Left = mirror(root.Right)
	newroot.Right = mirror(root.Left)
	return newroot
}

func compare(a,b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}else if a == nil{
		return false
	}else if b == nil{
		return false
	}
	if a.Val != b.Val {
		return false
	}else{
		return compare(a.Left,b.Left) && compare(a.Right,b.Right)
	}
}

//



func main() {
	fmt.Println("hello world")
}