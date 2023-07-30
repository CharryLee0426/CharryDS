package binaryindexsearch

import "testing"

func TestTemplate1(t *testing.T) {
	inputs := []struct{
		nums []int
		target int
		expected int
	}{
		{[]int{-1,0,3,5,9,12}, 9, 4},
		{[]int{-1,0,3,5,9,12}, 2, -1},
		{[]int{-1,0,3,5,9,12}, 12, 5},
		{[]int{-1,0,3,5,9,12}, -1, 0},
	}

	for _, input := range inputs {
		t1 := NewTemplate1(input.nums, input.target)
		actual := t1.Search()
		if actual != input.expected {
			t.Errorf("Test failed, expected: '%v', got:  '%v'", input.expected, actual)
		}
	}
}