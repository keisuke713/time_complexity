package main

import (
	"fmt"
)

func main() {
	var nums [5]int
	for i := 0; i < len(nums); i++ {
		fmt.Printf("index: %d, pointer: %p\n", i, &nums[i])
	}
}
