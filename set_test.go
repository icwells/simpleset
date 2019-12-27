// Tests simpleset package

package simpleset

import (
	"testing"
)

func evaluateLength(t *testing.T, a, e int) {
	// Compares actual length to expected
	if a != e {
		t.Errorf("Actual length %d does nto equal expected: %d", a, e)
	}
}

func evaluateBool(t *testing.T, a, e bool) {
	// Compares results of inset
	if a != e {
		t.Errorf("Actual InSet value %v does nto equal expected: %v", a, e)
	}
}

func TestSet(t *testing.T) {
	// Tests set attributes
	var cases [][]interface{}
	cases = append(cases, []interface{}{"a", "b", "c", "d", "e"})
	cases = append(cases, []interface{}{1, 2, 3, 4, 5})
	cases = append(cases, []interface{}{1.2, 2.58, 3.8, 4.6, 5.0})
	for idx, c := range cases {
		var s *Set
		switch idx {
		case 0:
			s = NewStringSet()
		case 1:
			s = NewIntSet()
		case 2:
			s = NewFloatSet()
		}
		evaluateLength(t, s.Length(), 0)
		for idx, i := range c {
			s.Add(i)
			evaluateLength(t, s.Length(), idx+1)
			evaluateBool(t, s.InSet(i), true)
		}
		l := s.Length()
		for _, i := range c {
			s.Pop(i)
			l--
			evaluateLength(t, s.Length(), l)
			evaluateBool(t, s.InSet(i), false)
		}
		s.Extend(cases[idx])
		evaluateLength(t, s.Length(), len(cases[idx]))
	}
}
