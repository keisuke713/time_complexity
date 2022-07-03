package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var nums1 [1000000]int

	// O(n)のコード
	// 配列のサイズが5->10で2倍になるとステップ数が5->10の2倍(n)
	iterateNumsOnce(nums1)
	iterateNumsOnce(nums2)

	// O(n^2)のコード
	// 配列のサイズが5->10で2倍になるとステップ数が25->100の4倍(n^2)
	iterateNumsTwice(nums1)
	iterateNumsTwice(nums2)
}

func iterateNumsOnce(nums []int) {
	for i, n := range nums {
		fmt.Printf("%d回目, 値: %d\n", i+1, n)
	}
}

func iterateNumsTwice(nums []int) {
	for i, n := range nums {
		for j, _ := range nums {
			fmt.Printf("%d回目, 値: %d\n", i*len(nums)+j+1, n)
		}
	}
}
