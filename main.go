package main

import (
	"collatz/collatz"
	"fmt"
)

func main() {
	n := 10000
	maxCount := 0
	maxN := 0

	for i := 0; i < n; i++ {
		tmp := executeCount(i)
		if maxCount < tmp {
			maxCount = tmp
			maxN = i
		}
	}

	fmt.Printf("maxN: %d, maxCount: %d", maxN, maxCount)
	return
}

func executeCount(n int) int {
	r := collatz.CountSteps(n)
	fmt.Printf("n: %d, count: %d\n", n, r)
	return r
}
