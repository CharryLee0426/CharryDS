package tree

import "testing"

func TestSegTree(t *testing.T) {
	input := []int{1, 3, 5}
	expectedOutput := []int{9, 8}

	root := &SegTree{0, len(input) - 1, nil, nil, 0}

	GenerateSegTree(root, input)

	output := []int{}
	output = append(output, QuerySegTree(root, 0, 2))
	UpdateSegTree(root, 1, 2)
	output = append(output, QuerySegTree(root, 0, 2))

	for i := 0; i < len(output); i++ {
		if output[i] != expectedOutput[i] {
			t.Errorf("Expected %d, got %d", expectedOutput[i], output[i])
		}
	}
}