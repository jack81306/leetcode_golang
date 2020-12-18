package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
    if len(digits) == 0{
		return []string{}
	}
	data := []string{"abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	res := []string{""}
	for i:=len(digits)-1;i>=0;i--{
		num := int(digits[i]-'2')
		tmp := []string{}
		for j:=0;j<len(res);j++{
			for k:=0;k<len(data[num]);k++{
				tmp = append(tmp,string(data[num][k])+res[j])
			}
		}
		res = tmp
	}
	return res
}

func main() {
	fmt.Println("hello world")
}