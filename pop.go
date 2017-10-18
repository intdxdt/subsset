package subsset

func (s *SubSSet) pop() interface{} {
	n := len(s.view) - 1
	val := s.view[n]

	s.view[n] = nil
	s.view = s.view[:n]
	s.j -= 1
	return val
}

func (s *SubSSet) popLeft() interface{} {
	val := s.view[0]
	s.view[0] = nil

	s.view = s.view[1:]
	s.i += 1
	return val
}
