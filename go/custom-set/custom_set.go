package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, elem := range l {
		set.Add(elem)
	}
	return set
}

func (s Set) String() string {
	var elements []string
	for elem := range s {
		elements = append(elements, fmt.Sprintf("%q", elem)) 
	}
	return "{" + strings.Join(elements, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
		if s2.Has(elem) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	intersection := New()
	for elem := range s1 {
		if s2.Has(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := New()
	for elem := range s1 {
		if !s2.Has(elem) {
			difference.Add(elem)
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	union := New()
	for elem := range s1 {
		union.Add(elem)
	}
	for elem := range s2 {
		union.Add(elem)
	}
	return union
}
