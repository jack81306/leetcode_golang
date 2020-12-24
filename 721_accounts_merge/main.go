package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	table := make([]int,len(accounts))
	key := []string{}
	for i,_ := range table{
		table[i]=i
	}
	hashmap := make(map[string][]int)
	for i,v := range accounts{
		for j:=1;j<len(v);j++{
			val,ok := hashmap[v[j]]
			if ok{
				hashmap[v[j]] = append(val,i)
			}else{
				key = append(key,v[j])
				hashmap[v[j]] = []int{i}
			}
		}
	}
	sort.Strings(key)
	for _,v := range hashmap{
		if len(v)>1{
			for i:=1;i<len(v);i++{
				accounti,accountj,groupi,groupj := v[0],v[i],table[v[0]],table[v[i]]
				for accounti!=groupi{
					accounti = groupi
					groupi = table[accounti]
				}
				for accountj!=groupj{
					accountj = groupj
					groupj = table[accountj]
				}
				if groupi!=groupj {
					tmpj := v[i]
					for tmpj != accountj{
						tmp := table[tmpj]
						table[tmpj] = groupi
						tmpj = tmp
					}
					table[tmpj] = groupi
				}
			}
		}
	}

	for i,v := range table{
		if i!=v{
			val := v		
			for table[val]!=val{
				val = table[val]
			}
			table[i] = val
		}
	}

	res := make([][]string,len(accounts))
	for i,_ := range res{
		res[i] = []string{accounts[i][0]}
	}

	for j:=0;j<len(key);j++{
		i,v := key[j],hashmap[key[j]]
		res[table[v[0]]] = append(res[table[v[0]]],i)
	}

	result := [][]string{}
	for _,v := range res{
		if len(v) > 1{
			result = append(result,v)
		}
	}

	return result
}



func main() {
	fmt.Println("hello world")
}

