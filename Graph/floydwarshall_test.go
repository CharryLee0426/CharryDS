package graph

import (
	"fmt"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	edges := [][]int{
		{0, 1, 2},
		{0, 2, 4},
		{1, 2, 1},
		{1, 3, 7},
		{2, 3, 3},
	}

	expectedOutput := [][]int {
		{0, 2, 3, 6},
		{2, 0, 1, 4},
		{3, 1, 0, 3},
		{6, 4, 3, 0},
	}

	fw := NewFloydWarshall(edges, 4)
	output := fw.Solve()

	fmt.Println(output)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if output[i][j] != expectedOutput[i][j] {
				t.Errorf("Expected %d, got %d", expectedOutput[i][j], output[i][j])
			}
		}
	}
}