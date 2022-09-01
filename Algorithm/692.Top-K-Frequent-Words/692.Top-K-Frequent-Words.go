package topkfrequentwords

import "container/heap"

type pair struct {
	word  string
	count int
}

type pq []pair

func (q pq) Len() int {
	return len(q)
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q pq) Less(i, j int) bool {
	s, t := q[i], q[j]
	return s.count < t.count || s.count == t.count && s.word > t.word
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.(pair))
}

func (q *pq) Pop() interface{} {
	a := *q
	v := a[len(a)-1]
	*q = a[:len(a)-1]
	return v
}

func topKFrequent(words []string, k int) []string {
	cnt := make(map[string]int)
	for _, word := range words {
		cnt[word]++
	}

	q := &pq{}
	for w, c := range cnt {
		heap.Push(q, pair{w, c})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(q).(pair).word
	}
	return ans
}
