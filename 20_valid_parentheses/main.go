package main

import (
	"fmt"
)


func push(s,str string)string{
	s = s+str
	return s
}
func pop(s string)string{
	s = s[:len(s)-1]
	return s
}
func peek(s string)byte{
	if len(s)==0{
		return '0'
	}
	return s[len(s)-1]
}

func isValid(s string) bool {
	stack := ""
	for _,v := range s{
		if v=='('|| v=='['|| v=='{'{
			stack = push(stack,string(v))
		}else if v==')'||v==']'||v=='}'{
			var target byte
			switch v{
				case ')':
					target = '('
				case ']':
					target = '['
				case '}':
					target = '{'
			}
			if target!=peek(stack){
				return false
			}else{
				stack = pop(stack)
			}
		}
	}
	if len(stack)!=0{
		return false
	}
	return true
}

func main() {
	fmt.Println("hello world")
}
