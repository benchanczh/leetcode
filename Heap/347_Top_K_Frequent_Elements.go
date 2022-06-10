package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	saved := map[int]int{}
	for _, num := range nums {
		if _, ok := saved[num]; ok {
			saved[num]++
		} else {
			saved[num] = 1
		}
	}

	h := &IntMapHeap{}
	result := []int{}
	for key, val := range saved {
		heap.Push(h, IntMap{key, val})
	}
	for i := 0; i < k; i++ {
		top := heap.Pop(h)
		result = append(result, top.(IntMap).Number)
	}
	return result
}

// An IntMapHeap is a min-heap of ints.
type IntMap struct {
	Number int
	Count  int
}
type IntMapHeap []IntMap

func (h IntMapHeap) Len() int {
	return len(h)
}

func (h IntMapHeap) Less(i, j int) bool {
	return h[i].Count > h[j].Count
}

func (h IntMapHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMapHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(IntMap))
}

func (h *IntMapHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 3
	fmt.Println(topKFrequent(nums, k))
}
