package rdyset

import (
	"testing"
)

func TestHas(t *testing.T) {
	s := Set(1, 2, 3)
	has := s.Has(4)
	print("Set has 2: ", has)
}

func TestAdd(t *testing.T) {
	s := Set(1, 2, 3)
	s.Add(4)

	got := s.String()
	expect := "{ 1, 2, 3, 4 }"

	if expect != got {
		t.Fatalf("Expected %s, got %s", expect, got)
	}
}

func TestDelete(t *testing.T) {
	s := Set(1, 2, 3, 4)
	s.Remove(4)
	expect := s.String()

	if expect != "{ 1, 2, 3 }" {
		t.Fail()
	}
}

