package collatz

func CountSteps(n int) int {
	if n <= 0 {
		return 0
	}

	c := 0
	for ; n != 1; c++ {
		n = executeStep(n)
	}
	return c
}

func executeStep(n int) int {
	if n == 1 {
		return n
	}
	if n%2 == 0 {
		return n / 2
	} else {
		return n*3 + 1
	}
}
