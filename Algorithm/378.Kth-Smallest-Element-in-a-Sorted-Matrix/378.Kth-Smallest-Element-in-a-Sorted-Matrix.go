package main

import (
	"container/heap"
	"fmt"
)

// Solution1: merge sort with heap
func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix)
	h := new(IHeap)
	for i := 0; i < n; i++ {
		h.Push(entry{matrix[i][0], i, 0})
	}
	for i := 0; i < k-1; i++ {
		e := heap.Pop(h).(entry)
		if e.col < n-1 {
			heap.Push(h, entry{matrix[e.row][e.col+1], e.row, e.col + 1})
		}
	}
	return heap.Pop(h).(entry).num
}

type entry struct {
	num, row, col int
}

type IHeap []entry

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i].num < h[j].num }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(v interface{}) {
	*h = append(*h, v.(entry))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := h.Len()
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// Solution2: binary search
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)>>1
		if check(matrix, mid, k, n) {
			// The target is in the left section, include the mid element.
			right = mid
		} else {
			// The target is in the right section, don't include the mid element.
			left = mid + 1
		}
	}
	// When left == right, the loop ended, so the left or right is the answer.
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	row, col := n-1, 0
	cnt := 0
	for row >= 0 && col < n {
		if matrix[row][col] <= mid {
			// The column above current row is smaller then mid
			cnt += row + 1
			// Go to the next column
			col++
		} else {
			// Go to the prev row
			row--
		}
	}
	return cnt >= k
}

func main() {
	fmt.Println(kthSmallest([][]int{{1, 2}, {1, 3}}, 2))
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
	fmt.Println(kthSmallest([][]int{{-5}}, 1))
}
