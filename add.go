package subsset

import "github.com/intdxdt/algor"

//Add - Push item to s - worst case at O(N^2)
//the cost here at O(n^2) is to allow dynamic indexing
//to add an item, Push uses O(lgN) to find where to insert
//and linear time O(1) or O(N-1) to keep s in sorted order
func (s *SubSSet) Add(item ...interface{}) *SubSSet {
	for _, v := range item {
		s.add(v)
	}
	return s
}

func (s *SubSSet) add(v interface{}) *SubSSet {
	if s.IsEmpty() {
		s.appendLeft(v)
		return s
	}
	//reserve enough room to the left and right
	s.reserve(true, true)
	array := s.base
	n := s.j - 1

	rawIdx := algor.BS(array, v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}

	//o := v.(int)
	if idx == s.i {
		// equal to i
		if rawIdx < 0 {
			s.appendLeft(nil) //make room
		}
		array[s.i] = v
	} else if s.i < idx && idx < n {
		//between i & j exclusive
		if rawIdx < 0 {
			s.append(nil) //make room
			copy(array[idx+1:n+2], array[idx:n+1])
		}
		array[idx] = v
	} else if idx == n {
		//equal to j-1
		if rawIdx < 0 {
			// >> array[n] >> array[n+1]
			s.append(array[n])
		}
		array[n] = v
	} else if idx > n {
		//greater than nth loc
		s.append(v)
	}
	return s
}
