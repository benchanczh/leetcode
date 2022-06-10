package main

func removeDuplicates(nums []int) int {
	j := 0
	// index j points to the last non-duplicated element
	// when i points to an element != nums[j], move j to next index and assign nums[i] to it
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
