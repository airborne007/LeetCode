package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Solution1: use buildin sort
func kClosest1(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

// Solution2: heap sort
func kClosest2(points [][]int, k int) [][]int {
	h := new(IHeap)
	heap.Init(h)
	for _, point := range points {
		x, y := point[0], point[1]
		heap.Push(h, [3]int{x, y, x*x + y*y})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([][]int, 0, k)
	for _, p := range *h {
		ans = append(ans, []int{p[0], p[1]})
	}
	return ans
}

type IHeap [][3]int

func (h IHeap) Len() int            { return len(h) }
func (h IHeap) Less(i, j int) bool  { return h[i][2] > h[j][2] }
func (h IHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(v interface{}) { *h = append(*h, v.([3]int)) }
func (h *IHeap) Pop() interface{}   { old := *h; n := h.Len(); v := old[n-1]; *h = old[:n-1]; return v }

// Solution3: quick select
func kClosest(points [][]int, k int) [][]int {
	rand.Seed(time.Now().UnixNano())
	n := len(points)
	arr := make([][3]int, n)
	for i, p := range points {
		arr[i] = [3]int{p[0], p[1], p[0]*p[0] + p[1]*p[1]}
	}
	var quickSelect func(int, int)
	quickSelect = func(left, right int) {
		pivot := rand.Int()%(right-left+1) + left
		arr[pivot], arr[right] = arr[right], arr[pivot]
		cnt := left
		for i := left; i < right; i++ {
			if arr[i][2] < arr[right][2] {
				arr[i], arr[cnt] = arr[cnt], arr[i]
				cnt++
			}
		}
		arr[cnt], arr[right] = arr[right], arr[cnt]
		if cnt+1 == k {
			return
		} else if cnt+1 < k {
			quickSelect(cnt+1, right)
		} else {
			quickSelect(left, cnt-1)
		}
	}
	quickSelect(0, n-1)
	ans := make([][]int, k)
	for i, p := range arr[:k] {
		ans[i] = []int{p[0], p[1]}
	}
	return ans
}

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
