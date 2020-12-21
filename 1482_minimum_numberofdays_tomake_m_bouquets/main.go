package main

import (
	"fmt"
)

func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay){
        return -1
    }
	var tmp int32 =1
	intmax,intmin := int(tmp<<31-1),int(tmp<<31)
	max,min := intmin,intmax
	for _,v := range bloomDay{
		if v > max {
			max = v
		}
		if v < min{
			min = v
		}
	}
	for max!=min{
		mid := int((max+min)/2)
		if check(bloomDay,k,mid)>=m{
			max = mid
		}else{
			min = mid+1
		}
	}
	return min
}

func check(bloomDay []int,k int,day int)int{
	cnt,res :=0,0
	for _,v := range bloomDay{
		if v <= day {
			cnt++
		}else{
			cnt=0
		}
		if cnt==k{
			res++
			cnt=0
		}
	}
	return res
}


func main() {
	fmt.Println("hello world")
}
