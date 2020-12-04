package main

import (
	"fmt"
)
//此為暴力解
//可用dfs解這題 能放左括號就放左括號，再放又掛號
func generateParenthesis(n int) []string {
	total := 1<<(n*2)
	now := ""
	cnt := 0 
	result := []string{}
	for i:=0;i<total;i++{
		now = ""
		cnt = 0
		tmp := i
		for j:=0;j<n*2;j++{
			if tmp&1 == 1 {
				cnt += 1
				now = now + "("
			}else{
				cnt -= 1
				now = now + ")"
			}
			if cnt < 0 {
				break
			}
			tmp = tmp>>1
		}
		if cnt == 0 {
			result = append(result,now)
		}
	}
	return result
}

func main()  {
	fmt.Println(generateParenthesis(3))
}