package main

import (
	"fmt"
)
//source :https://leetcode.com/problems/lru-cache/submissions/
type LRUCache struct {
    capacity    int
    currentTime int
    hash        map[int]int
    time        map[int]int
}

func Constructor(capacity int) LRUCache {
    // fmt.Println("Capacity is:", capacity)
    ret := LRUCache{}
    ret.capacity = capacity
    ret.hash = map[int]int{}
    ret.time = map[int]int{}
    ret.currentTime = 1
    return ret
}

func (this *LRUCache) Get(key int) int {
    this.currentTime += 1
    if _, ok := this.hash[key]; !ok {
        return -1
    }
    
    this.time[key] = this.currentTime
    return this.hash[key]
}

func (this *LRUCache) Put(key int, value int)  {
    this.currentTime += 1
    
    if _, ok := this.hash[key]; ok {
        this.hash[key] = value
        this.time[key] = this.currentTime
        return
    }
    
    if len(this.hash) == this.capacity {
        mn := -1
        lessUsed := 0

        for k, t := range this.time {
            if mn == -1 {
                mn = t
                lessUsed = k
                continue
            }
            if t < mn {
                lessUsed = k
                mn = t
            }
        }

        delete(this.time, lessUsed)
        delete(this.hash, lessUsed)
    }
    
    this.time[key] = this.currentTime
    this.hash[key] = value
    }
func main() {
	fmt.Println("hello world")
}
