package graph

import "testing"

func TestDijkstra(t *testing.T) {
	graph := [][]int{
		{1, 0, 1},
		{1, 2, 1},
		{2, 3, 1},
	}

	expected := map[int]int{
		0: 1,
		1: 0,
		2: 1,
		3: 2,
	}

	d := NewDijkstra(graph, 4)
	d.Solve(1)

	for i := 0; i < 4; i++ {
		if d.Min[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], d.Min[i])
		}
	}
}