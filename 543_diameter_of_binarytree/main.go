package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//recursive 同時求子樹最大深度與最大diameter

func diameterOfBinaryTree(root *TreeNode) int {
	v, _ := diametermax(root)
	return v
}

func diametermax(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l, ldepth := diametermax(root.Left)
	r, rdepth := diametermax(root.Right)
	maxdiameter := ldepth + rdepth
	if maxdiameter > l && maxdiameter > r {
		return maxdiameter, max(ldepth, rdepth) + 1
	} else {
		return max(l, r), max(ldepth, rdepth) + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println("hello world")
}
