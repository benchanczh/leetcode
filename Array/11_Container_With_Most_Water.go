package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := (right - left) * min(height[left], height[right])

	for left < right {
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
