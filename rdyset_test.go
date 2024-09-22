package rdyset

import (
	"fmt"
	"strings"
	"testing"
)

func TestHas(t *testing.T) {
	fmt.Println("Testing Has()...")
	var expect bool
	s := Set(1, 2, 3)
	
	expect = s.Has(1)
	if !expect {
		t.Fatalf("Fails expectation s.Has(1): %v", s.String())
	}

	expect = !s.Has(4)
	if !expect {
		t.Fatalf("Fails expectation !s.Has(4): %v", s.String())
	}
}

func TestAdd(t *testing.T) {
	fmt.Println("Testing Add()...")
	var expect bool
	s := Set(1, 2, 3)

	s.Add(4)

	expect = s.Equals(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation s.Equals(1, 2, 3, 4): %v", s.String())
	}

	expect = !s.Equals(Set(1, 2, 3)) && !s.Equals(Set(1, 2))
	if (!expect) {
		t.Fatalf("Fails expectation !s.Equals(Set(1, 2, 3)) && !s.Equals(Set(1, 2)): %v", s.String())
	}
}


func TestRemove(t *testing.T) {
	fmt.Println("Testing Remove()...")
	var expect bool
	s := Set(1, 2, 3, 4)
	s.Remove(4)

	expect = s.Equals(Set(1, 2, 3))
	if !expect {
		t.Fatalf("Fails expectation s.Equals(Set(1, 2, 3)): %v", s.String())
	}

	expect = !s.Equals(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation !s.Equals(Set(1, 2, 3, 4)): %v", s.String())
	}
}

func TestSize(t *testing.T) {
	fmt.Println("Testing Size()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Size() == 3
	if !expect {
		t.Fatalf("Fails expectation s.Size() == 3: %v", s.String())
	}

	s.Remove(3)

	expect = s.Size() == 2
	if !expect {
		t.Fatalf("Fails expectation s.Size() == 2: %v", s.String())
	}
}

func TestIsSubsetOf(t *testing.T) {
	fmt.Println("Testing IsSubsetOf()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsSubsetOf(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation s.IsSubsetOf(Set(1, 2, 3, 4)): %v", s.String())
	}

	expect = s.IsSubsetOf(Set(1, 2, 3))
	if !expect {
		t.Fatalf("Fails expectation s.IsSubsetOf(Set(1, 2, 3)): %v", s.String())
	}

	expect = s.IsSubsetOf(s)
	if !expect {
		t.Fatalf("Fails expectation s.IsSubsetOf(s): %v", s.String())
	}
}

func TestIsProperSubsetOf(t *testing.T) {
	fmt.Println("Testing IsProperSubsetOf()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsProperSubsetOf(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation s.IsProperSubsetOf(Set(1, 2, 3, 4)): %v", s.String())
	}

	expect = !s.IsProperSubsetOf(Set(1, 2, 3))
	if !expect {
		t.Fatalf("Fails expectation !s.IsProperSubsetOf(Set(1, 2, 3)): %v", s.String())
	}

	expect = !s.IsProperSubsetOf(s)
	if !expect {
		t.Fatalf("Fails expectation !s.IsProperSubsetOf(s): %v", s.String())
	}
}

func TestIsSupersetOf(t *testing.T) {
	fmt.Println("Testing IsSupersetOf()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsSupersetOf(Set(1, 2))
	if !expect {
		t.Fatalf("Fails expectation s.IsSupersetOf(Set(1, 2)): %v", s.String())
	}

	expect = !s.IsSupersetOf(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation !s.IsSupersetOf(Set(1, 2, 3, 4)): %v", s.String())
	}

	expect = s.IsSupersetOf(s)
	if !expect {
		t.Fatalf("Fails expectation s.IsSupersetOf(s): %v", s.String())
	}
}

func TestIsProperSupersetOf(t *testing.T) {
	fmt.Println("Testing IsProperSupersetOf()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.IsProperSupersetOf(Set(1, 2))
	if !expect {
		t.Fatalf("Fails expectation s.IsProperSupersetOf(Set(1, 2)): %v", s.String())
	}

	expect = !s.IsProperSupersetOf(Set(1, 2, 3))
	if !expect {
		t.Fatalf("Fails expectation !s.IsProperSupersetOf(Set(1, 2, 3)): %v", s.String())
	}

	expect = !s.IsProperSupersetOf(s)
	if !expect {
		t.Fatalf("Fails expectation !s.IsProperSupersetOf(s): %v", s.String())
	}
}

func TestEquals(t *testing.T) {
	fmt.Println("Testing Equals()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Equals(Set(1, 2, 3))
	if !expect {
		t.Fatalf("Fails expectation s.Equals(Set(1, 2, 3)): %v", s.String())
	}

	expect = s.Equals(s)
	if !expect {
		t.Fatalf("Fails expectation s.Equals(s): %v", s.String())
	}

	expect = !s.Equals(Set(1, 2)) 
	if !expect {
		t.Fatalf("Fails expectation !s.Equals(Set(1, 2)): %v", s.String())
	}

	s.Remove(3)

	expect = s.Equals(Set(1, 2))
	if !expect {
		t.Fatalf("Fails expectation s.Equals(Set(1, 2)): %v", s.String())
	}
}

func TestUnion(t *testing.T) {
	fmt.Println("Testing Union()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Union(Set(4)).Equals(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation s.Union(Set(4)).Equals(Set(1, 2, 3, 4)): %v", s.String())
	}

	expect = s.Union(Set(3, 4)).Equals(Set(1, 2, 3, 4))
	if !expect {
		t.Fatalf("Fails expectation s.Union(Set(4)).Equals(Set(1, 2, 3, 4)): %v", s.String())
	}

	expect = s.Union(Set[int]()).Equals(s)
	if !expect {
		t.Fatalf("Fails expectation s.Union(Set[int]()).Equals(s): %v", s.String())
	}
}

func TestIntersection(t *testing.T) {
	fmt.Println("Testing Intersection()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Intersection(Set(3, 4, 5)).Equals(Set(3))
	if !expect {
		t.Fatalf("Fails expectation s.Intersection(Set(3, 4, 5)).Equals(Set(3)): %v", s.String())
	}

	expect = s.Intersection(Set[int]()).Equals(Set[int]())
	if !expect {
		t.Fatalf("Fails expectation s.Intersection(Set[int]()).Equals(Set[int]()): %v", s.String())
	}
}

func TestDifference(t *testing.T) {
	fmt.Println("Testing Difference()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Difference(Set(3)).Equals(Set(1, 2))
	if !expect {
		t.Fatalf("Fails expectation s.Difference(Set(3)).Equals(Set(1, 2)): %v", s.String())
	}

	expect = s.Difference(Set(4, 5, 6)).Equals(s)
	if !expect {
		t.Fatalf("Fails expectation s.Difference(Set(4, 5, 6)).Equals(s): %v", s.String())
	}

	expect = s.Difference(Set(1, 2, 3, 4, 5, 6)).Equals(Set[int]())
	if !expect {
		t.Fatalf("Fails expectation s.Difference(Set(1, 2, 3, 4, 5, 6)).Equals(Set[int]()): %v", s.String())
	}

	expect = s.Difference(Set[int]()).Equals(s)
	if !expect {
		t.Fatalf("Fails expectation s.Difference(Set[int]()).Equals(s): %v", s.String())
	}
}

func TestSymmetricDifference(t *testing.T) {
	fmt.Println("Testing SymmetricDifference()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.SymmetricDifference(Set(4, 5, 6)).Equals(Set(1, 2, 3, 4, 5, 6))
	if !expect {
		t.Fatalf("Fails expectation s.SymmetricDifference(Set(4, 5, 6)).Equals(Set(1, 2, 3, 4, 5, 6)): %v", s.String())
	}

	expect = s.SymmetricDifference(Set(3, 4, 5)).Equals(Set(1, 2, 4, 5))
	if !expect {
		t.Fatalf("Fails expectation s.SymmetricDifference(Set(3, 4, 5)).Equals(Set(1, 2, 4, 5)): %v", s.String())
	}

	expect = s.SymmetricDifference(Set[int]()).Equals(s)
	if !expect {
		t.Fatalf("Fails expectation s.SymmetricDifference(Set[int]()).Equals(s): %v", s.String())
	}
}

func TestForEach(t *testing.T) {
	fmt.Println("Testing ForEach()...")
	s := Set(1, 2, 3)

	members := make(map[int]bool)
	s.ForEach(func(member int) {
		members[member] = true
	})

	for _, member := range []int{1, 2, 3} {
		if !members[member] {
			t.Fatalf("Fails expectation members[member] = true: %v, %v", members, member)
		}
	}
}

func TestFilter(t *testing.T) {
	fmt.Println("Testing Filter()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Filter(func(k int) bool {
		return k > 1
	}).Equals(Set(2, 3))
	if !expect {
		t.Fatalf("Fails expectation s.Filter(...).Equals(Set(2, 3)): %v", s.String())
	}
}

func TestMap(t *testing.T) {
	fmt.Println("Testing Map()...")
	var expect bool
	s := Set(1, 2, 3)

	expect = s.Map(func(k int) int {
		return k * k
	}).Equals(Set(1, 4, 9))
	if !expect {
		t.Fatalf("Fails expectation s.Map(...).Equals(Set(1, 4, 9)): %v", s.String())
	}
}

func TestArray(t *testing.T) {
	fmt.Println("Testing Array()...")
	s := Set(1, 2, 3)

	contains := func(s []int, e int) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}

		return false
	}

	members := []int{ 1, 2, 3 }
	s_array := s.Array()

	s.ForEach(func (member int) {
		if !contains(members, member) {
			t.Fatalf("Fails expectation !contains(members, member): %v, %v", s_array, member)
		}
	})
}

func TestString(t *testing.T) {
	fmt.Println("Testing String()...")
	var expect bool
	s := Set(1, 2, 3)

	str := s.String()

	s.ForEach(func (member int) {
		if !strings.Contains(str, fmt.Sprintf("%d", member)) {
			t.Fatalf("Fails expectation strings.Contains(str, fmt.Sprintf(...))): %v, %v", str, member)
		}
	})

	count := strings.Count(str, ",")
	expect = count == s.Size() - 1
	if !expect {
		t.Fatalf("Fails expectation count != len(s) - 1: %v, %v, %v", s, count, len(s) - 1)
	}
}