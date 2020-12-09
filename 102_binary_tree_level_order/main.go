package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	n     *TreeNode
	level int
}

//使用queue來取得level order
//使用channel代替queue
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make(chan node, 1024)
	queue <- node{n: root, level: 0}

	levelval := []int{}
	result := [][]int{}
	nowlevel := 0
	for len(queue) > 0 {
		treenode := <-queue
		if treenode.level != nowlevel {
			result = append(result, levelval)
			nowlevel = treenode.level
			levelval = []int{treenode.n.Val}
		} else {
			levelval = append(levelval, treenode.n.Val)
		}
		if treenode.n.Left != nil {
			queue <- node{n: treenode.n.Left, level: nowlevel + 1}
		}
		if treenode.n.Right != nil {
			queue <- node{n: treenode.n.Right, level: nowlevel + 1}
		}
	}
	result = append(result, levelval)
	return result
}

func main() {
	fmt.Println("hello world")
}
