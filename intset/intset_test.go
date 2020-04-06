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
