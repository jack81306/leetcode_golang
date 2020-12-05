package main

import (
	"fmt"
)

type Stack []int 

func (s Stack)push(i int)Stack{
	s = append(s,i)
	return s
}

func (s Stack)isempty() bool{
	return len(s)==0
}


func (s Stack)pop() Stack{
	s = s[:len(s)-1]
	return s
}

func (s Stack)peek() int{
	return s[len(s)-1]
}


//tle 暴力解
func dailyTemperaturesf(T []int) []int {
	if len(T) == 1 {
		return []int{0,}
	}
	cal := make([]int,len(T))
	result := make([]int,len(T))
	isget := make([]bool,len(T))
	for i:=1;i<len(T);i++{
		diff := T[i] - T[i-1]
		for j:=0;j<i;j++{
			cal[j] += diff
			if isget[j] == false&&cal[j] > 0 {
				result[j] = i-j
				isget[j] = true
			}
		}
	}
	return result
}

//stack 版本 建立一個stack當 
//stack為空或最頂值大於當前值:將當前值push進stack
//小於時:將stack最頂層值pop出來並計算，這代表棧最頂層的值遇到第一個大於他的直
//因此就可以計算結果，不斷這個工作直到棧為空或最頂大於當前值
func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0,}
	}
	result := make([]int,len(T))
	val := Stack{}
	index := Stack{}

	val = val.push(T[0])
	index = index.push(0)

	for i:=1;i<len(T);i++{
		fmt.Println(val)
		for {
			if val.isempty(){
				val = val.push(T[i])
				index = index.push(i)
				break
			}
			v := val.peek()
			fmt.Println(v,T[i])
			if v>=T[i]{
				val = val.push(T[i])
				index = index.push(i)
				break
			}else{
				val = val.pop()
				k := index.peek()
				index = index.pop()
				result[k] = i-k
			}
		}
	}
	return result
}

func main(){
	fmt.Println("hello world")
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}