package graph

import "testing"

func TestBellmanFord(t *testing.T) {
	graph := [][]int{
		[]int{0, 1, 100},
		[]int{1, 2, 100},
		[]int{0, 2, 500},
		[]int{2, 3, 100},
		[]int{0, 3, 200},
		[]int{3, 1, -150},
	}

	expected := []int{0, 50, 150, 200}

	b := NewBellmanFord(graph, 4)
	b.Solve(0)

	for i, v := range b.Min {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		} 
	}
}