package main

import (
	"fmt"
)

func main() {
	var nums []int
	for i := 0; i < 9; i++ {
		fmt.Println("================")
		nums = append(nums, i+1)
		// 2**n+1個目の要素を挿入するタイミングで配列を作り直している(2,3,5,9...)
		fmt.Printf("address: %p, len: %d, cap: %d \n", nums, len(nums), cap(nums))
		for i := 0; i < len(nums); i++ {
			fmt.Printf("index: %d, pointer: %p\n", i, &nums[i])
		}
	}
}
