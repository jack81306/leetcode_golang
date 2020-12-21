package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	start,end := 0,1
	scnt,tcnt := make([]int,52),make([]int,52)
	for _,c := range t{
		if c>='A' && c<='Z'{
			tcnt[c-'A'+26]++
		}else if c>='a' && c<='z'{
			tcnt[c-'a']++
		}
	}
	imin,jmin := -10000000,10000000
	isadd := 1
	for end<=len(s)&&end>start{
		if isadd>0{
			c := s[end-1]
			if c>='A' && c<='Z'{
				scnt[c-'A'+26]++
			}else if c>='a' && c<='z'{
				scnt[c-'a']++
			}
		}else{
			c := s[start-1]
			if c>='A' && c<='Z'{
				scnt[c-'A'+26]--
			}else if c>='a' && c<='z'{
				scnt[c-'a']--
			}
		}
		if check(scnt,tcnt){
			if end-start < jmin-imin{
				jmin = end
				imin = start
			}
			start++
			isadd = -1
		}else{
			end++
			isadd = 1
		}
	}
	if imin < 0{
		return ""
	}else{
		return s[imin:jmin]
	}
}
func check(s,t []int)bool{
	for i:=0;i<len(t);i++{
		if t[i]!=0&&t[i]>s[i]{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}