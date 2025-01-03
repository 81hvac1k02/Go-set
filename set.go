// Package set set.go
package set

type Set[T comparable] map[T]struct{}

// Add a new element to the set
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Clone creates a copy of the set.
func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T])
	for k := range s {
		clone.Add(k)
	}
	return clone
}

// Contains checks if an element is in the set
func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Difference creates a new set that is the difference of two sets.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			difference.Add(k)
		}
	}
	return difference
}

// Intersection creates a new set that is the intersection of two sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := make(Set[T])
	for k := range s {
		if other.Contains(k) {
			intersection.Add(k)
		}
	}
	return intersection
}

// IsSubset checks if one set is a subset of another.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// Map adds every value from a slice to the set.
func (s Set[T]) Map(values []T) {
	for _, value := range values {
		s.Add(value)
	}
}

// Remove an element from the set
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Size gets the size of the set
func (s Set[T]) Size() int {
	return len(s)
}

// Union creates a new set that is the union of two sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := s.Clone()
	for k := range other {
		union.Add(k)
	}
	return union
}
