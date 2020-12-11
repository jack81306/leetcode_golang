package main

import (
	"fmt"
)

type Trie struct {
	iskey bool
	next [26]*Trie
}

//prefixtree 詳解
//https://blog.csdn.net/lisonglisonglisong/article/details/45584721
/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this
    for i:=0;i<len(word);i++{
		idx := int(word[i]-'a')
		if cur.next[idx] == nil{
			t := new(Trie)
			cur.next[idx] = t
		}
		cur = cur.next[idx]
	}
	cur.iskey = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i:=0;i<len(word);i++{
		idx := int(word[i]-'a')
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}
	return cur.iskey
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    cur := this
	for i:=0;i<len(prefix);i++{
		idx := int(prefix[i]-'a')
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}
	return true
}




func main()  {
	fmt.Println()
}