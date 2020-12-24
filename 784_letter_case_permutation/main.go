package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {
    if len(S)==0{
		return []string{}
	}
	if len(S)==1{
		if S[0] >= '0' && S[0] <= '9'{
			return []string{S}
		}else{
			var base byte
			if S[0]>='a' && S[0]<='z'{
				base = S[0]-'a'
			}else{
				base = S[0]-'A'
			}
			return []string{string(base+'a'),string(base+'A')}
		}
	}
	tmp := letterCasePermutation(S[1:])
	res := []string{}
	if S[0] >= '0' && S[0] <= '9'{
		for _,v := range tmp{
			res = append(res,string(S[0])+v)
		}
	}else{
		var base byte
		if S[0]>='a' && S[0]<='z'{
			base = S[0]-'a'
		}else{
			base = S[0]-'A'
		}
		for _,v := range tmp{
			res = append(res,string(base+'a')+v)
			res = append(res,string(base+'A')+v)
		}
	}
	return res
}

func main() {
	fmt.Println("hello world")
}

