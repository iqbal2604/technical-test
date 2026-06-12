package main

import (
	"container/heap"
	"fmt"
)

// Min Heap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		// Kalau ukuran heap melebihi k, buang yang terkecil
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Root heap = elemen terkecil dari k terbesar = kth largest
	return (*h)[0]
}

func logic_prob_3_1() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("nums: %v, k: %d\n", nums, k)
	fmt.Printf("Kth largest: %d\n", findKthLargest(nums, k))
}
