package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
    n^=(n>>1)
    return n&(n+1)==0
}

func main() {
	fmt.Println("hello world")
}

