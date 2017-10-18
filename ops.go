package subsset

//Union - s union
func (s *SubSSet) Union(other *SubSSet) *SubSSet {
	u := s.Clone()
	for _, v := range other.view {
		u.Add(v)
	}
	return u
}

//Intersection - s intersection
func (s *SubSSet) Intersection(other *SubSSet) *SubSSet {
	inter := NewSubSSet(s.cmp)
	for _, v := range other.view {
		if s.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

//Difference- s difference
//items in s not contained in other
func (s *SubSSet) Difference(other *SubSSet) *SubSSet {
	diff := NewSubSSet(s.cmp)
	for _, v := range s.view {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

//SymDifference - symmetric difference with between s and other
//new s with elements in either s or other but not both
func (s *SubSSet) SymDifference(other *SubSSet) *SubSSet {
	return s.Difference(other).Union(
		other.Difference(s),
	)
}
