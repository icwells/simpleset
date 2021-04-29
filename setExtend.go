// Contains extend and fromSlice methods

package simpleset

// ExtendString adds values from a string slice to set.
func (s *Set) ExtendString(v []string) error {
	var err error
	if len(v) > 0 {
		err = s.checkType(v[0])
		if err == nil {
			for _, i := range v {
				s.s[i] = false
			}
		}
	}
	return err
}

// FromStringSlice returns a new set from a string slice.
func FromStringSlice(v []string) *Set {
	ret := NewStringSet()
	ret.ExtendString(v)
	return ret
}

// ExtendInt adds values from an int slice to set.
func (s *Set) ExtendInt(v []int) error {
	var err error
	if len(v) > 0 {
		err = s.checkType(v[0])
		if err == nil {
			for _, i := range v {
				s.i[i] = false
			}
		}
	}
	return err
}

// FromIntSlice returns a new set from a string slice.
func FromIntSlice(v []int) *Set {
	ret := NewIntSet()
	ret.ExtendInt(v)
	return ret
}

// ExtendFloat adds values from a float64 slice to set.
func (s *Set) ExtendFloat(v []float64) error {
	var err error
	if len(v) > 0 {
		err = s.checkType(v[0])
		if err == nil {
			for _, i := range v {
				s.f[i] = false
			}
		}
	}
	return err
}

// FromFloatSlice returns a new set from a float64 slice.
func FromFloatSlice(v []float64) *Set {
	ret := NewFloatSet()
	ret.ExtendFloat(v)
	return ret
}
