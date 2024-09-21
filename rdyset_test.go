package rdyset

import (
	"testing"
)

func fail(t *testing.T, value bool) {
	t.Fatalf("Expected %v; got %v", !value, value)
}

func TestHas(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)
	
	expect = s.Has(1)
	if !expect {
		fail(t, expect)
	}

	expect = !s.Has(4)
	if !expect {
		fail(t, expect)
	}
}

func TestAdd(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)
	s.Add(4)

	expect = s.Equals(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = !s.Equals(Set(1, 2, 3)) && !s.Equals(Set(1, 2))
	if (!expect) {
		fail(t, expect)
	}
}


func TestRemove(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3, 4)
	s.Remove(4)

	expect = s.Equals(Set(1, 2, 3))
	if !expect {
		fail(t, expect)
	}

	expect = !s.Equals(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}
}

func TestSize(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Size() == 3
	if !expect {
		fail(t, expect)
	}

	s.Remove(3)

	expect = s.Size() == 2
	if !expect {
		fail(t, expect)
	}
}

func TestIsSubsetOf(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsSubsetOf(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = s.IsSubsetOf(Set(1, 2, 3))
	if !expect {
		fail(t, expect)
	}

	expect = s.IsSubsetOf(s)
	if !expect {
		fail(t, expect)
	}
}

func TestIsProperSubsetOf(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsProperSubsetOf(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = !s.IsProperSubsetOf(Set(1, 2, 3))
	if !expect {
		fail(t, expect)
	}

	expect = !s.IsProperSubsetOf(s)
	if !expect {
		fail(t, expect)
	}
}

func TestIsSupersetOf(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsSupersetOf(Set(1, 2))
	if !expect {
		fail(t, expect)
	}

	expect = !s.IsSupersetOf(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = s.IsSupersetOf(s)
	if !expect {
		fail(t, expect)
	}
}

func TestIsProperSupersetOf(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsProperSupersetOf(Set(1, 2))
	if !expect {
		fail(t, expect)
	}

	expect = !s.IsProperSupersetOf(Set(1, 2, 3))
	if !expect {
		fail(t, expect)
	}

	expect = !s.IsProperSupersetOf(s)
	if !expect {
		fail(t, expect)
	}
}

func TestEquals(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Equals(Set(1, 2, 3))
	if !expect {
		fail(t, expect)
	}

	expect = s.Equals(s)
	if !expect {
		fail(t, expect)
	}

	expect = !s.Equals(Set(1, 2)) 
	if !expect {
		fail(t, expect)
	}

	s.Remove(3)

	expect = s.Equals(Set(1, 2))
	if !expect {
		fail(t, expect)
	}
}

func TestUnion(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Union(Set(4)).Equals(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = s.Union(Set(3, 4)).Equals(Set(1, 2, 3, 4))
	if !expect {
		fail(t, expect)
	}

	expect = s.Union(Set[int]()).Equals(s)
	if !expect {
		fail(t, expect)
	}
}

func TestIntersection(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Intersection(Set(3, 4, 5)).Equals(Set(3))
	if !expect {
		fail(t, expect)
	}

	expect = s.Intersection(Set[int]()).Equals(Set[int]())
	if !expect {
		fail(t, expect)
	}
}

func TestDifference(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Difference(Set(3)).Equals(Set(1, 2))
	if !expect {
		fail(t, expect)
	}

	expect = s.Difference(Set(4, 5, 6)).Equals(s)
	if !expect {
		fail(t, expect)
	}

	expect = s.Difference(Set(1, 2, 3, 4, 5, 6)).Equals(Set[int]())
	if !expect {
		fail(t, expect)
	}

	expect = s.Difference(Set[int]()).Equals(s)
	if !expect {
		fail(t, expect)
	}
}

func TestSymmetricDifference(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.SymmetricDifference(Set(4, 5, 6)).Equals(Set(1, 2, 3, 4, 5, 6))
	if !expect {
		fail(t, expect)
	}

	expect = s.SymmetricDifference(Set(3, 4, 5)).Equals(Set(1, 2, 4, 5))
	if !expect {
		fail(t, expect)
	}

	expect = s.SymmetricDifference(Set[int]()).Equals(s)
	if !expect {
		fail(t, expect)
	}
}

func TestForEach(t *testing.T) {
	// TODO
}

func TestFilter(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Filter(func(k int) bool {
		return k > 1
	}).Equals(Set(2, 3))
	if !expect {
		fail(t, expect)
	}
}

func TestMap(t *testing.T) {
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Map(func(k int) int {
		return k * k
	}).Equals(Set(1, 4, 9))
	if !expect {
		fail(t, expect)
	}
}

func TestArray(t *testing.T) {
	// TODO
}

func TestString(t *testing.T) {
	// TODO
}