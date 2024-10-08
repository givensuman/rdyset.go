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
	s.Add(members...)

	return s
}

/*
Adds a new member to the set

Ensures a ∈ A
 */
func (A set[T]) Add(members ...T) {
	for _, member := range members {
		A[member] = struct{}{}
	}
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
The equality of two sets

Set A is equal to set B if A is a subset of B, 
and B is a subset of A

Returns A = B
 */
func (A set[T]) Equals(B set[T]) bool {
	return A.IsSubsetOf(B) && B.IsSubsetOf(A)
}

/*
The union of two sets

The set C is the union of sets A and B if it contains
all the members of A and B

Returns A ∪ B
 */
func (A set[T]) Union(B set[T]) set[T] {
	C := Set[T]()

	A.ForEach(func(member T) {
		C.Add(member)
	})

	B.ForEach(func(member T) {
		C.Add(member)
	})

	return C
}

/*
The intersection of two sets

The set C is the intersection of sets A and B
if it contains only the members of A which are also
in B, and only the members of B which are also in A

Returns A ∩ B
 */
func (A set[T]) Intersection(B set[T]) set[T] {
	C := Set[T]()

	if A.Size() > B.Size() {
		B.ForEach(func(member T) {
			if A.Has(member) {
				C.Add(member)
			}
		})
	} else {
		A.ForEach(func(member T) {
			if B.Has(member) {
				C.Add(member)
			}
		})
	}

	return C
}

/*
The difference of two sets

The set C is the difference of B from A (e.g. A - B)
if it contains all the members of A which are not in B

Returns A ∖ B
 */
func (A set[T]) Difference(B set[T]) set[T] {
	C := Set[T]()

	A.ForEach(func(member T) {
		if !B.Has(member) {
			C.Add(member)
		}
	})

	return C
}

/*
The symmetric difference of two sets

The set C is the symmetric difference of A and B
if it contains all the members of A which are not in B,
and all the members of B which are not in A

Returns A Δ B
 */
func (A set[T]) SymmetricDifference(B set[T]) set[T] {
	C := A.Difference(B)
	D := B.Difference(A)
	return C.Union(D)
}

/*
Iterates over the unordered set members
and calls the provided function
 */
func (A set[T]) ForEach(do func(member T)) {
	for member := range A {
		do(member)
	}
}

/*
Returns a set of members that meet
condition using provided function
 */
func (A set[T]) Filter(do func(member T) bool) set[T] {
	C := Set[T]()

	A.ForEach(func(a T) {
		if do(a) {
			C.Add(a)
		}
	})

	return C
}

/*
Converts the set to a slice of its members
 */
func (s set[T]) Slice() []T {
	var members []T
	for member := range s {
		members = append(members, member)
	}

	return members
}

/*
Converts the set to a string of the form { a, b, c }
 */
func (s set[t]) String() string {
	var members []string
	for key := range s {
		members = append(members, fmt.Sprintf("%v", key))
	}
	return "{ " + strings.Join(members, ", ") + " }"
}
