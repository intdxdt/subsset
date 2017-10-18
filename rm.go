package subsset

import "github.com/intdxdt/algor"

func (s *SubSSet) Remove(item ...interface{}) *SubSSet {
	for _, v := range item {
		s.rm(v)
	}
	return s
}

//Remove an item by value from the Array
func (s *SubSSet) rm(v interface{}) *SubSSet {
	if s.IsEmpty() {
		return s
	}
	array := s.base
	n := s.j - 1

	rawIdx := algor.BS(array, v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}

	if rawIdx >= 0 {
		if idx == s.i {
			// equal to i
			s.PopLeft()
		} else if s.i < idx && idx < n {
			//between i & j exclusive
			copy(array[idx:n+1], array[idx+1:n+1])
			s.Pop()
		} else if idx == n {
			//equal to nth item
			s.Pop()
		}
	}
	return s
}
