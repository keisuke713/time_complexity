package main

import "fmt"

func main() {
	var arr [1]int
	for i, n := range arr {
		for _, n2 := range arr {
			fmt.Printf("index: %d, n: %d", i, n*n2)
		}
	}
}

// func
