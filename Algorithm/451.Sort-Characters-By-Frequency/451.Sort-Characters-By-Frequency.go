package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"sort"
)

// Solution1: bucket sort
func frequencySort(s string) string {
	freq := make(map[byte]int)
	maxFreq := 0
	for i := range s {
		freq[s[i]]++
		maxFreq = max(maxFreq, freq[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, cnt := range freq {
		buckets[cnt] = append(buckets[cnt], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i >= 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Solution2: quick sort
func frequencySort2(s string) string {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}
	pairs := make([]pair, 0, len(freq))
	for ch, cnt := range freq {
		pairs = append(pairs, pair{ch, cnt})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })
	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}

// Solution3: Heap sort
func frequencySort3(s string) string {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}
	h := new(IHeap)
	for ch, cnt := range freq {
		heap.Push(h, pair{ch, cnt})
	}
	ans := make([]byte, 0, len(s))
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}

type pair struct {
	ch  byte
	cnt int
}

type IHeap []pair

func (h IHeap) Len() int            { return len(h) }
func (h IHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h IHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *IHeap) Pop() interface{}   { old := *h; n := h.Len(); v := old[n-1]; *h = old[:n-1]; return v }

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
