package rdyset

import (
	"fmt"
	"strings"
)

type set[T comparable] map[T]struct{}

/*
Instantiates a new Set 
*/
func Set[T comparable](members ...T) set[T] {
	s := make(set[T])
	for _, member := range members {
		s.Add(member)
	}

	return s
}

/*
Adds a new member to the set

Ensures a ∈ A
*/
func (A set[T]) Add(member T) {
	A[member] = struct{}{}
}

/*
Removes an existing member from the set

Ensures a ∉ A
*/
func (A set[T]) Remove(member T) {
	delete(A, member)
}

/*
Whether an element if a member of the set

Returns a ∈ A
*/
func (A set[T]) Has(member T) bool {
	_, exists := A[member]
	return exists
}

/*
The number of members of the set, also called its cardinality

Returns |A|
*/
func (A set[T]) Size() int {
	return len(A)
}

/*
Whether or not the set is a subset of argument

Set A is a subset of B if every element of
A is contained in B

Returns A ⊆ B
*/
func (A set[T]) IsSubsetOf(B set[T]) bool {
	for member := range A {
		if !B.Has(member) {
			return false
		}
	}

	return true
}

/*
Whether or not the set is a proper subset of argument

Set A is a proper subset of B if every element of 
A is contained in B, and B has at least one element
not contained within A

Returns A ⊂ B
*/
func (A set[T]) IsProperSubsetOf(B set[T]) bool {
	return A.IsSubsetOf(B) && B.Size() > A.Size()
}

/*
Whether or not the set is a superset of argument

Set A is a superset of B if every element of
B is contained within B

Returns A ⊇ B
*/
func (A set[T]) IsSupersetOf(B set[T]) bool {
	return B.IsSubsetOf(A)
}

/*
Whether or not the set is a proper superset of argument

Set A is a proper superset of B if every element of
B is contained within A, and A has at least one element
not contained within B

Returns A ⊃ B
*/
func (A set[T]) IsProperSupersetOf(B set[T]) bool {
	return B.IsProperSubsetOf(A)
}

/*
Whether or not two sets are equal

Set A is equal to set B if they contain all of 
each other's members and are of the same size

Returns A = B
*/
func (A set[T]) Equals(B set[T]) bool {
	return A.IsProperSubsetOf(B) && B.Size() == A.Size()
}

func (A set[T]) Union(B set[T]) set[T] {
	C := Set[T]()

	A.ForEach(func(member T) {
		C.Add(member)
	})
	B.ForEach(func(member T)) {
		C.Add(member)
	}

	return C
}

func (A set[T]) ForEach(do func(member T)) {
	for member := range A {
		do(member)
	}
}

/*
Converts the set to an unordered list
*/
func (s set[T]) Array() []T {
	var members []T
	for member := range s {
		members = append(members, member)
	}

	return members
}

/*
Converts the set to its string representation
*/
func (s set[t]) String() string {
	var members []string
	for key, _ := range s {
		members = append(members, fmt.Sprintf("%v", key))
	}
	return "{ " + strings.Join(members, ", ") + " }"
}
