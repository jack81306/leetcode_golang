package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	lo,hi := 0,0
	for _,v := range nums{
		lo = (lo^v) & ^hi
		hi = (hi^v) & ^lo
	}

	return lo
}

func main() {
	fmt.Println("hello world")
}
