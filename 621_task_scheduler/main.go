package main

import (
	"fmt"
)

//greedy
//更優解:https://leetcode.com/problems/task-scheduler/discuss/466959/Python-Easiest-Solution
func leastInterval2(tasks []byte, n int) int {
	alltask := make([]int, 26)
	cooldown := make([]int, 26)
	for _, v := range tasks {
		i := int(v - 'A')
		alltask[i]++
	}

	result := 0
	for !allzero(alltask) {
		t := picktask(alltask, cooldown)
		if t == -1 {
			result++
			minusone(cooldown)
		} else {
			result++
			alltask[t]--
			minusone(cooldown)
			cooldown[t] = n
		}
	}
	return result
}

func minusone(all []int) {
	for i := range all {
		if all[i] > 0 {
			all[i] = all[i] - 1
		}
	}
}

func allzero(all []int) bool {
	for _, v := range all {
		if v != 0 {
			return false
		}
	}
	return true
}

func picktask(all []int, cooldown []int) int {
	max := -1
	maxcnt := 0
	for idx, val := range all {
		if cooldown[idx] == 0 && val > 0 && val > maxcnt {
			max = idx
			maxcnt = val
		}
	}
	return max
}

func main() {
	fmt.Println("hello world")
}
