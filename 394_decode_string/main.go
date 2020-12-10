package main

import (
	"fmt"
)

type Stack []byte

//用STACK來解決多個括號:以字元處理
//討論區使用字串處理
func decodeString(s string) string {
	stack := Stack{}
	nums := []int{}
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= "9" {
			cnt *= 10
			cnt += int(s[i] - '0')
			continue
		}
		if cnt != 0 {
			nums = append(nums, cnt)
			cnt = 0
		}
		if s[i] != ']' {
			stack = stack.push(s[i])
		} else {
			result := ""
			for {
				c := stack.peek()
				stack = stack.pop()
				if c == '[' {
					break
				}
				result = string(c) + result
			}
			s2 := decode(result, nums[len(nums)-1])
			nums = nums[:len(nums)-1]
			for j := 0; j < len(s2); j++ {
				stack = stack.push(s2[j])
			}
		}
	}
	return string(stack)
}

func decode(s string, cnt int) string {
	res := ""
	for i := 0; i < cnt; i++ {
		res += s
	}
	return res
}

func (s Stack) push(b byte) Stack {
	s = append(s, b)
	return s
}

func (s Stack) pop() Stack {
	if len(s) == 0 {
		return s
	}
	s = s[:len(s)-1]
	return s
}

func (s Stack) peek() byte {
	if len(s) == 0 {
		return '-'
	}
	return s[len(s)-1]
}

func main() {
	fmt.Println("hello world")
}
