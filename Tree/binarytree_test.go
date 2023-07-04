package tree

import "testing"

func TestPreOrder(t *testing.T) {
	inputs := []int{2, 5, 1, 3, 7}
	expected := []int{2, 5, 3, 7, 1}
	root := NewFromArray(inputs)

	res := root.PreOrder()

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], res[i])
		}
	}
}

func TestInorder(t *testing.T) {
	inputs := []int{2, 5, 1, 3, 7}
	expected := []int{3, 5, 7, 2, 1}

	root := NewFromArray(inputs)

	res := root.Inorder()

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], res[i])
		}
	}
}

func TestPostorder(t *testing.T) {
	inputs := []int{2, 5, 1, 3, 7}
	expected := []int{3, 7, 5, 1, 2}

	root := NewFromArray(inputs)

	res := root.Postorder()

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], res[i])
		}
	}
}

func TestLevelorder(t *testing.T) {
	inputs := []int{2, 5, 1, 3, 7}
	expected := []int{2, 5, 1, 3, 7}

	root := NewFromArray(inputs)

	res := root.Levelorder()

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], res[i])
		}
	}
}

func TestCheckSameTree(t *testing.T) {
	inputs := []struct {
		treeA []int
		treeB []int
		isSame bool
	} {
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
	}

	for _, input := range inputs {
		treeA := NewFromArray(input.treeA)
		treeB := NewFromArray(input.treeB)

		res := treeA.CheckSameTree(treeB)

		if res != input.isSame {
			t.Errorf("Expected %v, got %v", input.isSame, res)
		}
	}
}

func TestCheckFather(t *testing.T) {
	inputs := []struct {
		treeA []int
		treeB []int
		isFather bool
	} {
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3}, []int{2, 3}, false},
		{[]int{1, 2, 3}, []int{3}, true},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2, 3, 1, 3, 4, 3, 2}, []int{1, 2}, true},
	}

	for i, input := range inputs {
		treeA := NewFromArray(input.treeA)
		treeB := NewFromArray(input.treeB)

		res := treeA.CheckFather(treeB)

		if res != input.isFather {
			t.Errorf("round %v Expected %v, got %v\n", i, input.isFather, res)
		}
	}
}