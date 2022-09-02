package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Solution1: use small Heap, time O(nlogk), space O(logk)
func findKthLargest1(nums []int, k int) int {
	h := new(IHeap)
	heap.Init(h)
	for _, num := range nums {
		h.push(num)
		if h.Len() > k {
			h.pop()
		}
	}
	return h.pop()
}

type IHeap struct {
	sort.IntSlice
}

func (h *IHeap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *IHeap) Pop() interface{} {
	n := h.Len()
	v := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return v
}

func (h *IHeap) push(v int) {
	heap.Push(h, v)
}

func (h *IHeap) pop() int {
	return heap.Pop(h).(int)
}

// Solution2: quick select, time O(n), space O(logn)
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(arr []int, left, right, index int) int {
	q := partition(arr, left, right)
	if q == index {
		return arr[q]
	} else if q < index {
		return quickSelect(arr, q+1, right, index)
	} else {
		return quickSelect(arr, left, q-1, index)
	}
}

func partition(arr []int, left, right int) int {
	pivot := rand.Int()%(right-left+1) + left
	num := arr[pivot]
	arr[pivot], arr[right] = arr[right], arr[pivot]
	l := left - 1
	for i := left; i < right; i++ {
		if arr[i] < num {
			l++
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
	arr[l+1], arr[right] = arr[right], arr[l+1]
	return l + 1
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
