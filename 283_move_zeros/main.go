package main

import (
	"fmt"
)

//紀錄要往前swap幾個
//另解:紀錄當前應該放到哪 0跳過並最後補上
func moveZeroes(nums []int)  {
	zerocnt:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==0 {
			zerocnt++
			continue
		}
		swap(nums,i,i-zerocnt)
	}
}

func swap(nums []int,i int,j int){
	if i==j {
		return
	}
	nums[i] = nums[j]^nums[i]
	nums[j] = nums[i]^nums[j]
	nums[i] = nums[i]^nums[j]
}

func main()  {
	fmt.Println("hello world")
}