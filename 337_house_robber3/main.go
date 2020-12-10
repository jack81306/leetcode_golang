package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(robb(root, true), robb(root, false))
}

//recursion 1st version
func robb(root *TreeNode, robroot bool) int {
	if root == nil {
		return 0
	}
	if robroot {
		v := robb(root.Left, false) + robb(root.Right, false) + root.Val
		return v
	} else {
		l := max(robb(root.Left, true), robb(root.Left, false))
		r := max(robb(root.Right, true), robb(root.Right, false))
		return l + r
	}
}

func robc(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	l := robc(root.Left)
	r := robc(root.Right)

	return []int{l[0] + r[0] + root.Val, max(l[0], l[1]) + max(r[0], r[1])}

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
