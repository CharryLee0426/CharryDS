package graph

import "testing"

func TestPrim(t *testing.T) {
	input := [][]int{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}

	expected := 19

	p := NewPrim(input, 4)

	if p.Solve() != expected {
		t.Errorf("Expected %d, got %d", expected, p.Solve())
	}
}