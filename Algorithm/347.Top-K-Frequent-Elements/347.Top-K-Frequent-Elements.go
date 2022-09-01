package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	h := new(IHeap)
	heap.Init(h)
	for key, val := range freq {
		heap.Push(h, entry{num: key, freq: val})
		for h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(entry).num)
	}
	return ans
}

type entry struct {
	num, freq int
}

type IHeap []entry

func (h IHeap) Len() int            { return len(h) }
func (h IHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h *IHeap) Push(v interface{}) { *h = append(*h, v.(entry)) }
func (h *IHeap) Pop() interface{} {
	old := *h
	n := h.Len()
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
