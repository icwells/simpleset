// Package simpleset python-style set for string, int, and float64 types
package simpleset

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

// Set defines sets for string, int, and float64.
type Set struct {
	t     rune
	s     map[string]bool
	i     map[int]bool
	f     map[float64]bool
	types map[rune]string
}

// NewStringSet returns pointer to empty string set.
func NewStringSet() *Set {
	var s Set
	s.t = 's'
	s.s = make(map[string]bool)
	s.types = map[rune]string{'s': "string", 'i': "int", 'f': "float64"}
	return &s
}

// NewIntSet returns pointer to empty integer set.
func NewIntSet() *Set {
	var s Set
	s.t = 'i'
	s.i = make(map[int]bool)
	s.types = map[rune]string{'s': "string", 'i': "int", 'f': "float64"}
	return &s
}

// NewFloatSet returns pointer to empty float64 set.
func NewFloatSet() *Set {
	var s Set
	s.t = 'f'
	s.f = make(map[float64]bool)
	s.types = map[rune]string{'s': "string", 'i': "int", 'f': "float64"}
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
		err = errors.New("Slice must be strings, integers, or float64")
	}
	for _, i := range v {
		ret.Add(i)
	}
	return ret, err
}

// Clear removes all entries from set.
func (s *Set) Clear() {
	switch s.t {
	case 's':
		s.s = make(map[string]bool)
	case 'i':
		s.i = make(map[int]bool)
	case 'f':
		s.f = make(map[float64]bool)
	}
}

// ToStringSlice returns sorted slice of strings from set. Converts integers and floats to string if needed.
func (s *Set) ToStringSlice() []string {
	var ret []string
	switch s.t {
	case 's':
		for k := range s.s {
			ret = append(ret, k)
		}
	case 'i':
		for k := range s.i {
			ret = append(ret, strconv.Itoa(k))
		}
	case 'f':
		for k := range s.f {
			ret = append(ret, strconv.FormatFloat(k, 'f', -1, 64))
		}
	}
	sort.Strings(ret)
	return ret
}

// ToIntSlice returns sorted slice of integers from set.
func (s *Set) ToIntSlice() ([]int, error) {
	var ret []int
	var err error
	if s.t != 'i' {
		err = fmt.Errorf("Cannot convert %s set to int slice", s.types[s.t])
	} else {
		for k := range s.i {
			ret = append(ret, k)
		}
		sort.Ints(ret)
	}
	return ret, err
}

// ToFloatSlice returns sorted slice of integers from set.
func (s *Set) ToFloatSlice() ([]float64, error) {
	var ret []float64
	var err error
	if s.t != 'f' {
		err = fmt.Errorf("Cannot convert %s set to float slice", s.types[s.t])
	} else {
		for k := range s.f {
			ret = append(ret, k)
		}
		sort.Float64s(ret)
	}
	return ret, err
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

// Returns type as rune and an error if it does not equal s.t.
func (s *Set) checkType(v interface{}) error {
	var t rune
	var err error
	switch v.(type) {
	case string:
		t = 's'
	case int:
		t = 'i'
	case float64:
		t = 'f'
	default:
		err = errors.New("Value must be string, integer, or float64")
	}
	if err == nil && t != s.t {
		err = fmt.Errorf("%s submitted to %s set", s.types[t], s.types[s.t])
	}
	return err
}

// Add adds new value to set. Reutrns an error if v is of the wrong type.
func (s *Set) Add(v interface{}) error {
	err := s.checkType(v)
	if err == nil {
		switch s.t {
		case 's':
			s.s[v.(string)] = false
		case 'i':
			s.i[v.(int)] = false
		case 'f':
			s.f[v.(float64)] = false
		}
	}
	return err
}

// Extend adds all elements of slice to set.
func (s *Set) Extend(v []interface{}) error {
	var err error
	for _, i := range v {
		err = s.Add(i)
		if err != nil {
			break
		}
	}
	return err
}

// InSet returns true if value is in set.
func (s *Set) InSet(v interface{}) (bool, error) {
	var ret bool
	err := s.checkType(v)
	if err == nil {
		switch s.t {
		case 's':
			_, ret = s.s[v.(string)]
		case 'i':
			_, ret = s.i[v.(int)]
		case 'f':
			_, ret = s.f[v.(float64)]
		}
	}
	return ret, err
}

// Pop removes value from set. Returns an error if v is not a string, int, or float.
func (s *Set) Pop(v interface{}) error {
	var err error
	ex, err := s.InSet(v)
	if err == nil && ex {
		switch s.t {
		case 's':
			delete(s.s, v.(string))
		case 'i':
			delete(s.i, v.(int))
		case 'f':
			delete(s.f, v.(float64))
		}
	}
	return err
}
