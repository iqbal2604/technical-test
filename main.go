// package main

// import "fmt"

// var grid = [][]int{
// 	{0, 0, 1},
// 	{0, 1, 1},
// 	{0, 0, 0},
// 	{1, 0, 0},
// 	{0, 0, 0},
// 	{0, 1, 0},
// 	{0, 1, 1},
// 	{0, 0, 1},
// 	{1, 0, 0},
// }

// func dfs(grid [][]int, row, col int) {
// 	rows := len(grid)
// 	cols := len(grid[0])

// 	// Base case: out of bounds atau laut
// 	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 0 {
// 		return
// 	}

// 	// Tandai sudah dikunjungi
// 	grid[row][col] = 0

// 	// Kunjungi 4 arah
// 	dfs(grid, row+1, col)
// 	dfs(grid, row-1, col)
// 	dfs(grid, row, col+1)
// 	dfs(grid, row, col-1)
// }

// func countIslands(grid [][]int) int {
// 	count := 0

// 	for row := 0; row < len(grid); row++ {
// 		for col := 0; col < len(grid[0]); col++ {
// 			if grid[row][col] == 1 {
// 				count++
// 				dfs(grid, row, col)
// 			}
// 		}
// 	}

// 	return count
// }

// func main() {
// 	fmt.Println("Grid:")
// 	for _, row := range grid {
// 		fmt.Println(row)
// 	}

// 	result := countIslands(grid)
// 	fmt.Printf("\nJumlah pulau: %d\n", result)
// }

// package main

// import "fmt"

// func main() {
// 	n := 10
// 	a := 0
// 	b := 1

// 	if n <= 0 {
// 		fmt.Println("Jumlah suku harus lebih dari 0.")
// 		return
// 	}

// 	for i := 1; i < n; i++ {
// 		fmt.Printf("%d", a)
// 		if i < n-1 {
// 			fmt.Printf(", ")
// 		}
// 		next := a + b
// 		a = b
// 		b = next
// 	}
// 	fmt.Println()
// }

// package main

// import (
// 	"container/heap"
// 	"fmt"
// )

// // Min Heap implementation
// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func findKthLargest(nums []int, k int) int {
// 	h := &MinHeap{}
// 	heap.Init(h)

// 	for _, num := range nums {
// 		heap.Push(h, num)
// 		// Kalau ukuran heap melebihi k, buang yang terkecil
// 		if h.Len() > k {
// 			heap.Pop(h)
// 		}
// 	}

// 	// Root heap = elemen terkecil dari k terbesar = kth largest
// 	return (*h)[0]
// }

// func main() {
// 	nums := []int{3, 2, 1, 5, 6, 4}
// 	k := 2
// 	fmt.Printf("nums: %v, k: %d\n", nums, k)
// 	fmt.Printf("Kth largest: %d\n", findKthLargest(nums, k))
// }
