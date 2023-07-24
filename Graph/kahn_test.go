package graph

import "testing"

func TestKahn(t *testing.T) {
	graph := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
	}

	n := 4

	k := NewKahn(graph, n)

	route := k.Solve()

	expected1 := []int{0, 1, 2, 3}
	expected2 := []int{0, 2, 1, 3}

	if !equal(route, expected1) && !equal(route, expected2) {
		t.Errorf("Expected: %v or %v, got: %v", expected1, expected2, route)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}