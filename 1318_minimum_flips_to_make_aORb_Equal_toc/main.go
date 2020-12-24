package main

import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	cnt := 0
	tmp := c-(a|b&c)
	//a or b is zero but c is one
	for tmp>0{
		cnt++
		tmp&=tmp-1
	}
	//b or a is one but c is zero
	tmp = ((a|b)^c)-(c-(a|b&c))
	tmp2 := tmp&a
	for tmp2>0{
		cnt++
		tmp2&=tmp2-1
	}
	tmp2 = tmp&b
	for tmp2>0{
		cnt++
		tmp2&=tmp2-1
	}
	return cnt
}

func main() {
	fmt.Println("hello world")
}

