package main

import (
	"collatz/collatz"
	"fmt"
)

func main() {
	executeCount(10)
	return
}

func executeCount(n int) {
	fmt.Printf("n = %d: %d", n, collatz.CountSteps(n))
}
