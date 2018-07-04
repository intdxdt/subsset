package subsset

import (
	"fmt"
	"bytes"
	"github.com/intdxdt/math"
	"github.com/intdxdt/algor"
)

const N = 32

//SSet type
type SubSSet struct {
	cmp      func(a, b interface{}) int
	base     []interface{}
	view     []interface{}
	i        int
	j        int
	initSize int
}

//New Sorted Set
func NewSubSSet(cmp func(a, b interface{}) int, initSize ...int) *SubSSet {
	var iSize = N
	if len(initSize) > 0 {
		iSize = math.MaxInt(1, initSize[0])
	}

	var base, view, i, j = initQue(iSize)
	return &SubSSet{
		cmp:      cmp,
		base:     base,
		view:     view,
		i:        i,
		j:        j,
		initSize: iSize,
	}
}

//reveal underlying sorted slice of data view
func (s *SubSSet) DataView() []interface{} {
	return s.view
}

//Clone SSet
func (s *SubSSet) Clone() *SubSSet {
	var base = make([]interface{}, len(s.base))
	copy(base, s.base)
	var view = base[s.i:s.j]
	return &SubSSet{
		cmp:  s.cmp,
		base: base,
		view: view,
		i:    s.i,
		j:    s.j,
	}
}

//Contains item for the presence of a value in the Array - O(lgN)
func (s *SubSSet) Contains(items ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}
	var bln = true
	var n = len(items)
	for i := 0; bln && i < n; i++ {
		bln = algor.BS(s.base, items[i], s.cmp, s.i, s.j-1) >= 0
	}
	return bln
}

//IndexOf item in the sorted s  - O(lgN)
func (s *SubSSet) IndexOf(item interface{}) int {
	var idx = -1
	if s.IsEmpty() {
		return idx
	}
	idx = algor.BS(s.view, item, s.cmp, 0, s.len()-1)
	if idx < 0 {
		idx = -1
	}
	return idx
}

//Size of list
func (s *SubSSet) Size() int {
	return s.len()
}

//First item in s
func (s *SubSSet) First() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.first()
	}
	return r
}

//Last Item in s
func (s *SubSSet) Last() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.last()
	}
	return r
}

//NextItem given item in the sorted s
func (s *SubSSet) NextItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	var array = s.base
	var n = s.j - 1
	var rawIdx = algor.BS(array[:], v, s.cmp, s.i, n)

	var idx = rawIdx
	if idx < 0 {
		idx = -idx - 1
	}
	var next interface{}
	if rawIdx >= 0 && s.i <= idx && idx < n {
		next = array[idx+1]
	}
	return next
}

//PrevItem gets previous given item in the sorted s
func (s *SubSSet) PrevItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	var array = s.base
	var n = s.j - 1
	var rawIdx = algor.BS(array[:], v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}
	var prev interface{}
	if rawIdx >= 0 && s.i < idx && idx <= n {
		prev = array[idx-1]
	}
	return prev
}

//Filters items based on predicate : func (item Item, i int) bool
func (s *SubSSet) Filter(fn func(interface{}, int) bool) []interface{} {
	var items = make([]interface{}, 0)
	s.ForEach(func(v interface{}, i int) bool {
		if fn(v, i) {
			items = append(items, v)
		}
		return true
	})
	return items
}

//Pop item from the end of the sorted list
func (s *SubSSet) Pop() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.pop()
	}
	return r
}

//PopLeft item from the beginning of the sorted list
func (s *SubSSet) PopLeft() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.popLeft()
	}
	return r
}

//Values in s
func (s *SubSSet) Values() []interface{} {
	var values = make([]interface{}, s.len())
	copy(values, s.view)
	return values
}

//Empty SubSSet
func (s *SubSSet) Empty() *SubSSet {
	s.clear()
	return s
}

//Extend SSet given list of values as params
func (s *SubSSet) Extend(values ...interface{}) *SubSSet {
	for _, v := range values {
		s.Add(v)
	}
	return s
}

//First value in SSet
func (s *SubSSet) Get(idx int) interface{} {
	if idx < 0 {
		idx += len(s.view)
	}
	return s.view[idx]
}

//Checks if SSet empty
func (s *SubSSet) IsEmpty() bool {
	return s.len() == 0
}

func (s *SubSSet) String() string {
	var buffer bytes.Buffer
	var n = s.len()
	var token string
	buffer.WriteString("[")
	for i, o := range s.view {
		token = fmt.Sprintf("%v", o)
		if i < n-1 {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

//Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (s *SubSSet) ForEach(fn func(interface{}, int) bool) {
	for i, v := range s.view {
		if !fn(v, i) {
			break
		}
	}
}
