package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}


type deque []int

func push(s deque,x int)deque{
	return append(s,x)
}
func pop(s deque)deque{
	return s[:len(s)-1]
}
func peek(s deque)int{
	return s[len(s)-1]
}
func lpop(s deque)deque{
	return s[1:len(s)]
}
func lpeek(s deque)int{
	return s[0]
}
