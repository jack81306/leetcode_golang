package main

import (
	"fmt"
	"strconv"
	"strings"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return topreorder(root) 
}

func topreorder(root *TreeNode) string{
	if root == nil {
		return "X "
	}
	s := strconv.Itoa(root.Val)
	l := topreorder(root.Left)
	r := topreorder(root.Right)
	return s + " " + l + r
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode { 
	if len(data) == 0 {
		return nil
	}
	s := strings.Fields(data)
	root,_ := ds(s,0)
	return root
}

func ds(preorder []string,index int)(*TreeNode,int){
	if preorder[index] == "X"{
		return nil,index+1
	}
	s,_ := strconv.Atoi(preorder[index])
	root := new(TreeNode)
	root.Val = s
	var newidx int
	root.Left , newidx = ds(preorder,index+1)
	root.Right , newidx = ds(preorder,newidx)
	return root,newidx
}


func main()  {
	fmt.Println("hello world")
}