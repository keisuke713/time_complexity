package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1024
	numberOfStep := 0
	// O(n)とかだとn個のインプットがあって1,2,3..nと一つずつ処理を順番に行われるが、O(log n)は一回のステップで要素数が1/2になるためこのような計算量になる。
	for n > 1 {
		n = int(math.Floor(float64((n / 2))))
		numberOfStep++
		fmt.Printf("n: %d, numberOfStep: %d \n", n, numberOfStep)
	}
}
