package main

import (
	"fmt"
)

func main() {
	var nums []int8
	for i := 0; i < 10; i++ {
		nums = append(nums, int8(i+1))
	}
	// メモリは1バイトごと増えている
	for i := 0; i < len(nums); i++ {
		fmt.Printf("index: %d, address: %p\n", i, &nums[i])
	}
}
