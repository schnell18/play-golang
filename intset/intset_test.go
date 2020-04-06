package intset

import "testing"

func TestHas(t *testing.T) {
	ints := []int{1, 10, 100}
	s1 := IntSet{}

	for _, i := range ints {
		s1.Add(i)
	}

	for _, i := range ints {
		if !s1.Has(i) {
			t.Errorf("IntSet %v does not contain %d", s1, i)
		}
	}
}

func TestAdd(t *testing.T) {
	ints := []int{1, 10, 100}
	s1 := IntSet{}

	for _, i := range ints {
		s1.Add(i)
	}

	for _, i := range ints {
		if !s1.Has(i) {
			t.Errorf("IntSet %v does not contain %d", s1, i)
		}
	}
	s1.Add(300)
	if !s1.Has(300) {
		t.Errorf("IntSet %v does not contain 300", s1)
	}
}

func TestUnionWith(t *testing.T) {
	seta := []int{1, 10, 100}
	setb := []int{10, 11, 101}
	s1 := IntSet{}
	s2 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	for _, i := range setb {
		s2.Add(i)
	}

	s1.UnionWith(&s2)

	if !s1.Has(1) {
		t.Errorf("IntSet %v does not contain 1", s1)
	}
	if !s1.Has(10) {
		t.Errorf("IntSet %v does not contain 10", s1)
	}
	if !s1.Has(11) {
		t.Errorf("IntSet %v does not contain 11", s1)
	}
	if !s1.Has(100) {
		t.Errorf("IntSet %v does not contain 100", s1)
	}
	if !s1.Has(101) {
		t.Errorf("IntSet %v does not contain 101", s1)
	}
}

func TestString(t *testing.T) {
	exp := "{1 10 11 100 101}"
	seta := []int{1, 11, 10, 101, 100}
	s1 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	if s1.String() != exp {
		t.Errorf("IntSet %v s1.String() returns %v != %s", s1, &s1, exp)
	}

}

func TestLen(t *testing.T) {
	seta := []int{1, 11, 10, 101, 100}
	s1 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	if s1.Len() != 5 {
		t.Errorf("IntSet %v s1.Len() returns %v != 5", s1, s1.Len())
	}

}

func TestRemove(t *testing.T) {
	exp := "{1 10 11 100 101}"
	seta := []int{1, 11, 10, 101, 100}
	s1 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	// should take no effect
	s1.Remove(55)
	if s1.String() != exp {
		t.Errorf("IntSet %v s1.remove(55) returns %v != %s", s1, &s1, exp)
	}

	// should delete 101
	s1.Remove(101)
	exp = "{1 10 11 100}"
	if s1.String() != exp {
		t.Errorf("IntSet %v s1.remove(101) returns %v != %s", s1, &s1, exp)
	}

}

func TestClear(t *testing.T) {
	exp := "{}"
	seta := []int{1, 11, 10, 101, 100}
	s1 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	// should be empty
	s1.Clear()
	if s1.String() != exp {
		t.Errorf("IntSet s1.Clear() returns %v != %s", &s1, exp)
	}

}

func TestCopy(t *testing.T) {
	exp := "{1 10 11 100 101}"
	seta := []int{1, 11, 10, 101, 100}
	s1 := IntSet{}

	for _, i := range seta {
		s1.Add(i)
	}

	s2 := s1.Copy()
	if s2.String() != exp {
		t.Errorf("Cloned IntSet %v s2.String() returns %v != %s", *s2, s2, exp)
	}

	// add 55 to the clone and the original copy should remove unchanged
	s2.Add(55)
	exp2 := "{1 10 11 55 100 101}"
	if s2.String() != exp2 {
		t.Errorf("Cloned IntSet %v calls s2.Add(55) returns %v != %s", s2, &s1, exp2)
	}
	if s1.String() != exp {
		t.Errorf("Original IntSet %v s1.String() returns %v != %s", s1, s1.String(), exp)
	}
}
