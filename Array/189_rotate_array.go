package main

func rotateArray(nums []int, k int) {
	k %= len(nums)
	reverse(0, len(nums)-1, nums)
	reverse(0, k-1, nums)
	reverse(k, len(nums)-1, nums)
}

func reverse(a, b int, nums []int) {
	for a < b {
		nums[a], nums[b] = nums[b], nums[a]
		a++
		b--
	}
}
