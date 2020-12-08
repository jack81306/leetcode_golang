package main

import (
	"fmt"
	"sort"
    "strings"
)
//將字串排序後使用map結構儲存
func groupAnagrams(strs []string) [][]string {
	allresult := make(map[string][]string)

	for _,s := range strs{
		s2 := SortString(s)
		v , ok := allresult[s2]
		if ok {
			allresult[s2] = append(v,s)
		}else{
			allresult[s2] = []string{s,}
		}
	}

	result := [][]string{}
	for _,val := range allresult {
		result = append(result,val)
	}
	return result
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}


func main()  {
	fmt.Println("hello world")
}