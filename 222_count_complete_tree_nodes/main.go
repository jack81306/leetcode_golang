package main

import (
	"fmt"
)

 type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
 }
//iterative
func countNodes2(root *TreeNode) int {
	if root==nil{
		return 0
	}
	r,minlayer := root,0
	for r!=nil{
		minlayer++
		r = r.Right
	}
	left,right := 1<<(minlayer-1),1<<minlayer-1
	for left!=right{
		mid := int((left+right)/2)
		r = root
		for i:=minlayer-2;i>=0;i--{
			if (mid>>i)&1 == 1{
				r = r.Right
			}else{
				r = r.Left
			}
		}
		if r.Left==nil && r.Right==nil{
			right = mid
		}else if r.Left!=nil && r.Right!=nil{
			left = mid+1
		}
	}
	
	r = root
	for i:=minlayer-2;i>=0;i--{
		if (left>>i)&1 == 1{
			r = r.Right
		}else{
			r = r.Left
		}
	}
	if r.Left!=nil&&r.Right!=nil{
		return 2*left+1
	}else if r.Left==nil && r.Right==nil{
		return 2*left-1
	}
	return 2*left
}
//recursion
func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	ld,rd := depth(root.Left),depth(root.Right)
	if ld==rd{
		return 1<<ld+countNodes(root.Right)
	}else{
		return 1<<rd+countNodes(root.Left)
	}
}

func depth(root *TreeNode)int{
	if root==nil{
		return 0
	}
	return depth(root.Left)+1
}

func main() {
	fmt.Println("hello world")
}