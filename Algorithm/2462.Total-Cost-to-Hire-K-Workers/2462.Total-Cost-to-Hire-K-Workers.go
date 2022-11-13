package main

import "container/heap"

type pair struct {
	index, cost int
}

type Heap []pair

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}
	return h[i].cost < h[j].cost
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *Heap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func totalCost(costs []int, k int, candidates int) int64 {
	ans, n := 0, len(costs)
	h := &Heap{}
	var left, right int
	for i := 0; i < candidates; i++ {
		heap.Push(h, pair{i, costs[i]})
	}
	var j int
	for j = n - 1; j >= n-candidates && j >= candidates; j-- {
		heap.Push(h, pair{j, costs[j]})
	}

	left, right = candidates-1, j+1
	for ; k > 0; k-- {
		p := heap.Pop(h).(pair)
		index, cost := p.index, p.cost
		ans += cost
		if index <= left && left < right-1 {
			left++
			heap.Push(h, pair{left, costs[left]})
		} else if index >= right && left < right-1 {
			right--
			heap.Push(h, pair{right, costs[right]})
		}
	}
	return int64(ans)
}
