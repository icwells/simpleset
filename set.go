// Package simpleset python-style set for string, int, and float64 types
package simpleset

import (
	"errors"
	"sort"
)

// Set defines set for string, int, and float64.
type Set struct {
	t rune
	s map[string]bool
	i map[int]bool
	f map[float64]bool
}

// NewStringSet returns pointer to empty string set
func NewStringSet() *Set {
	var s Set
	s.t = 's'
	s.s = make(map[string]bool)
	return &s
}

// NewIntSet returns pointer to empty integer set
func NewIntSet() *Set {
	var s Set
	s.t = 'i'
	s.i = make(map[int]bool)
	return &s
}

// NewFloatSet returns pointer to empty float64 set
func NewFloatSet() *Set {
	var s Set
	s.t = 'f'
	s.f = make(map[float64]bool)
	return &s
}

// ToSet converts slice of strings, integers, or float64 to apprpriate set.
func ToSet(v []interface{}) (*Set, error) {
	var ret *Set
	var err error
	switch v[0].(type) {
	case string:
		ret = NewStringSet()
	case int:
		ret = NewIntSet()
	case float64:
		ret = NewFloatSet()
	default:
		err = errors.New("Slice must be strings, integers, or float64.")
	}
	for _, i := range v {
		ret.Add(i)
	}
	return ret, err
}

// ToStringSlice returns sorted slice of strings from set.
func (s *Set) ToStringSlice() []string {
	var ret []string
	for k := range s.s {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

// ToIntSlice returns sorted slice of integers from set.
func (s *Set) ToIntSlice() []int {
	var ret []int
	for k := range s.i {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

// ToFloatSlice returns sorted slice of integers from set.
func (s *Set) ToFloatSlice() []float64 {
	var ret []float64
	for k := range s.f {
		ret = append(ret, k)
	}
	sort.Float64s(ret)
	return ret
}

// Length returns length of set.
func (s *Set) Length() int {
	var ret int
	switch s.t {
	case 's':
		ret = len(s.s)
	case 'i':
		ret = len(s.i)
	case 'f':
		ret = len(s.f)
	}
	return ret
}

// Add adds new value to set.
func (s *Set) Add(v interface{}) {
	switch s.t {
	case 's':
		s.s[v.(string)] = false
	case 'i':
		s.i[v.(int)] = false
	case 'f':
		s.f[v.(float64)] = false
	}
}

// Extend adds all elements of slice to set.
func (s *Set) Extend(v []interface{}) {
	for _, i := range v {
		s.Add(i)
	}
}

// InSet returns true if value is in set.
func (s *Set) InSet(v interface{}) bool {
	var ret bool
	switch s.t {
	case 's':
		_, ret = s.s[v.(string)]
	case 'i':
		_, ret = s.i[v.(int)]
	case 'f':
		_, ret = s.f[v.(float64)]
	}
	return ret
}

// Pop removes value from set.
func (s *Set) Pop(v interface{}) {
	if s.InSet(v) {
		switch s.t {
		case 's':
			delete(s.s, v.(string))
		case 'i':
			delete(s.i, v.(int))
		case 'f':
			delete(s.f, v.(float64))
		}
	}
}
