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
//更優解:https://leetcode.com/problems/path-sum-iii/discuss/91878/17-ms-O(n)-java-Prefix-sum-method
func pathSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}
	queue := make(chan *TreeNode,5000)
	queue <- root
	result := 0
	for len(queue) > 0 {
		n := <- queue
		if n == nil {
			continue
		}
		result += ps(n,sum)
		queue<-n.Left
		queue<-n.Right
	}
	return result
}

func ps(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	s := ps(root.Left,sum-root.Val) + ps(root.Right,sum-root.Val)
	if root.Val == sum {
		s += 1
	}
	return s
}



func main() {
	fmt.Println("hello world")
}