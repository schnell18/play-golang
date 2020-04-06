package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports if the set contains the given non-negative integer x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the given non-negative integer x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Len returns the length of IntSet
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				len++
			}
		}
	}
	return len
}

// Remove deletes the give non-negative integer from the IntSet
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

// Clear empties the IntSet
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy makes a new copy of the IntSet
func (s *IntSet) Copy() *IntSet {
	t := IntSet{}
	t.words = make([]uint64, len(s.words))
	for i, sword := range s.words {
		t.words[i] = sword
	}
	return &t
}

// UnionWith calculates and sets s to s union t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, tword)
		} else {
			s.words[i] |= tword
		}
	}
}

// String create string representation of the IntSet
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
