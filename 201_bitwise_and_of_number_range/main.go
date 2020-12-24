package main

import (
	"fmt"
)


func rangeBitwiseAnd(m int, n int) int {
	cnt :=0
	for i:=30;i>=0;i--{
		if (1<<i)&m != (1<<i)&n {
			return cnt
		}else{
			if (1<<i)&m>0{
				cnt+=(1<<i)
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(1<<30)
}

