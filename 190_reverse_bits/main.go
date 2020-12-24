package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	//0~15 16~31
	for i:=0;i<16;i++{
		lo := (1<<(15-i)&num)<<(i*2+1)//low bit value
		hi := (1<<(16+i)&num)>>(i*2-1)//high bit value 
		num = num&(^(1<<(15-i)))|hi
		num = num&(^(1<<(16+i)))|lo
	}
	return num
}

func main() {
	fmt.Println("hello world")
}

