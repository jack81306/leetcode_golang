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

type queue struct{
	ch chan int
}

func getqueue(buf int)*queue{
	q := new(queue)
	q.ch = make(chan int,buf)
	return q
}

func (q *queue)enqueue(x int){
	q.ch <- x
}
func (q *queue)dequeue()int{
	return <-q.ch
}
func (q queue)len()int{
	return len(q.ch)
}

func max(a ...int)int{
	m := a[0]
	for _,v := range a{
		if v > m{
			m = v
		}
	}
	return m
}
