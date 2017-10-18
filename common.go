package subsset

func initQue(initSize int) ([]interface{}, []interface{}, int, int) {
	i := initSize / 2
	j := i
	base := make([]interface{}, initSize, initSize)
	view := base[i:j]
	return base, view, i, j
}

//Length of number of items in SSet
func (s *SubSSet) len() int {
	return len(s.view)
}

//First value in SSet
func (s *SubSSet) first() interface{} {
	return s.Get(0)
}

//Last value in SSet
func (s *SubSSet) last() interface{} {
	return s.Get(-1)
}

//Append to right side of SSet
func (s *SubSSet) append(o interface{}) *SubSSet {
	s.reserve(false, true)
	s.base[s.j] = o
	s.j += 1
	s.view = s.base[s.i: s.j]
	return s
}

//AppendLeft: appends to left of SSet
func (s *SubSSet) appendLeft(o interface{}) *SubSSet {
	s.reserve(true, false)

	if s.atPivot() {
		s.j += 1
	} else {
		s.i -= 1
	}
	s.base[s.i] = o

	s.view = s.base[s.i: s.j]
	return s
}

//reserve enough space left or right
// sufficient to contain elements on insert
func (s *SubSSet) reserve(left, right bool) {
	if left && s.i == 0 {
		s.expandBase()
	}

	if right && s.j == len(s.base) {
		s.expandBase()
	}
}

func (s *SubSSet) expandBase() {
	bn := len(s.base)
	vn := len(s.view)

	nn := 2 * bn
	if vn+(nn/2-vn/2) >= nn {
		nn = 2 * nn //not big enough
	}

	k := nn / 2
	mid := vn / 2

	ii := k - mid
	jj := ii + vn

	newBase := make([]interface{}, nn)
	copy(newBase[k:], s.view[mid:])
	copy(newBase[k-mid:k], s.view[0:mid])
	s.base = newBase

	s.i, s.j = ii, jj
	s.view = s.base[s.i: s.j]
}

func (s *SubSSet) atPivot() bool {
	n := len(s.base)
	return s.i == s.j && (s.i >= 0 && s.i < n)
}

//Clear everything in SSet
func (s *SubSSet) clear() *SubSSet {
	s.base, s.view, s.i, s.j = initQue(s.initSize)
	return s
}

