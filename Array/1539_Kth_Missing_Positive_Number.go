package main

func findKthPositive(arr []int, k int) int {
	a := 0
	for _, num := range arr {
		diff := num - a
		if diff > 1 {
			result := a + min(k, diff-1)
			k -= diff - 1
			if k <= 0 {
				return result
			}
		}
		a = num
	}
	// If k still > 0
	return arr[len(arr)-1] + k
}
