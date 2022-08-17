package designanorderedstream

type OrderedStream struct {
	data []string
	ptr  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data: make([]string, n),
	}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.data[idKey-1] = value
	start := s.ptr
	for s.ptr < len(s.data) && s.data[s.ptr] != "" {
		s.ptr++
	}
	return s.data[start:s.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
