package graph

import "testing"

func TestKruskal(t *testing.T) {
	input := [][]int{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}

	expected := [][]int{
		{2, 3, 4},
		{0, 3, 5},
		{0, 1, 10},
	}

	ks := NewKruskal(input, 4)
	ks.Solve()

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if ks.Tree[i][j] != expected[i][j] {
				t.Errorf("Expected %d, got %d", expected[i][j], ks.Tree[i][j])
			}
		}
	}
}