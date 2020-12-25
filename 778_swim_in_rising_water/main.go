package main

import (
	"fmt"
)
type pair struct{
	x,y int
}


func swimInWater(grid [][]int) int {
	l,r := 0,len(grid)*len(grid)-1
	for l!=r{
        //fmt.Println(l,r)
		mid := int((l+r)/2)
		if bfs(grid,mid){
			r=mid
		}else{
			l=mid+1
		}
	}
	return l
}
func bfs(grid [][]int,water int)bool{
	table := make([][]bool,len(grid))
	for i,_ := range table{
		table[i] = make([]bool,len(grid))
	}
	q := make(chan pair,3000)
	if water >= grid[0][0]{
		q <- pair{0,0}
	}
	for len(q)>0{
		cnt := len(q)
		for i:=0;i<cnt;i++{
			p := <-q
			if p.x<0||p.x>=len(grid)||p.y<0||p.y>=len(grid)||table[p.x][p.y]||water<grid[p.x][p.y]{
				continue
			}
			table[p.x][p.y]=true
			if p.x == len(grid)-1 && p.y==len(grid)-1{
                return true
			}
			q<-pair{p.x+1,p.y}
			q<-pair{p.x-1,p.y}
			q<-pair{p.x,p.y+1}
			q<-pair{p.x,p.y-1}
		}
	}
	return false
}


func main() {
	fmt.Println("hello world")
}