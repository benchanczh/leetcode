package main

func climbStairs(n int) int {
	a, b := 1, 1
	// Calculate until n == 2 since we already know b == 1 when n == 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return b
}
