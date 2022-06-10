package main

func largestRectangleArea(heights []int) int {
	result := 0
	stack := []int{-1}
	heights = append(heights, 0)
	n := len(heights)

	for i := 0; i < n; i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			result = max(result, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
