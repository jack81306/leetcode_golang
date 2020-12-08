package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//inorder 取得順序
func kthSmallest(root *TreeNode, k int) int {
    r := inorder(root)
    return r[k-1]
}

func inorder(root *TreeNode)[]int{
    if root == nil {
        return []int{}
    }
    l := inorder(root.Left)
    r := inorder(root.Right)
    result := []int{root.Val,}
    result = append(l,result...)
    result = append(result,r...)
    return result
}

func main()  {
	fmt.Println("hello world")
}