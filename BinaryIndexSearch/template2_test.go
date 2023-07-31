package binaryindexsearch

import "testing"

func TestTemplate2(t *testing.T) {
	inputs := []struct{
		nums []int
		target int
		expected int}{
			{[]int{1, 2, 3, 4, 5}, 3, 2},
		}

	for _, input := range inputs {
		actual := NewTemplate2(input.nums, input.target).Search()
		if actual != input.expected {
			t.Errorf("Got %d, expected %d", actual, input.expected)
		}
	}
}