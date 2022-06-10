package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := []int{}

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		}
	}
	return result
}

func intersect2(nums1 []int, nums2 []int) []int {
	saved := make(map[int]int)
	result := []int{}

	for _, num := range nums1 {
		if _, ok := saved[num]; ok {
			saved[num]++
		} else {
			saved[num] = 1
		}
	}

	for _, num := range nums2 {
		if val, ok := saved[num]; ok && val > 0 {
			result = append(result, num)
			saved[num]--
		}
	}

	return result
}
