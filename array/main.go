package main

import (
	"fmt"
)

func main() {
	var nums []int8
	for i := 0; i < 10; i++ {
		// fmt.Println("================")
		nums = append(nums, int8(i+1))
		// 2**n+1個目の要素を挿入するタイミングで配列を作り直している(2,3,5,9...)
		// fmt.Printf("len: %d, cap: %d \n", len(nums), cap(nums))
		// for i := 0; i < len(nums); i++ {
		// 	fmt.Printf("index: %d, pointer: %p\n", i, &nums[i])
		// }
	}
	// メモリは1バイトごと増えている
	for i := 0; i < len(nums); i++ {
		fmt.Printf("index: %d, address: %p\n", i, &nums[i])
	}
}
