package heap

import "testing"

func TestInsert(t *testing.T) {
	inputs := []Element{1, 2, 3, 4, 5, 6, 7}

	h := New(7)
	for _, input := range inputs {
		h.Insert(input)
	}

	// check the heap
	for i := 1; i <= 7; i++ {
		if h.elements[i] != Element(i) {
			t.Errorf("heap.elements[%d] = %d, want %d", i, h.elements[i], i)
		}
	}
}

func TestPeek(t *testing.T) {
	inputs := []Element{7, 6, 5, 4, 3, 2, 1}

	h := New(7)
	for _, input := range inputs {
		h.Insert(input)
	}

	if h.Peek() != 1 {
		t.Errorf("h.Peek() = %d, want %d", h.Peek(), 1)
	}
}

func TestPop(t *testing.T) {
	inputs := []Element{7, 5, 6, 2, 3, 1, 4}
	expected := []Element{1, 2, 3, 4, 5, 6, 7}

	h := New(7)
	for _, input := range inputs {
		h.Insert(input)
	}

	for i := 1; i <= 7; i++ {
		if h.Pop() != Element(expected[i-1]) {
			t.Errorf("h.Pop() = %d, want %d", h.Pop(), i)
		}
	}
}

func TestHeaptify(t *testing.T) {
	inputs := []Element{7, 5, 6, 2, 3, 1, 4}
	expected := []Element{1, 2, 3, 4, 5, 6, 7}

	h := NewFromArray(inputs, 7)

	for i := 1; i <= 7; i++ {
		if h.Pop() != Element(expected[i-1]) {
			t.Errorf("h.Pop() = %d, want %d", h.Pop(), i)
		}
	}
}