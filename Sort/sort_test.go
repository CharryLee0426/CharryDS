package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.BubbleSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("BubbleSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestSelectionSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.SelectionSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("SelectionSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestInsertionSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.InsertionSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("InsertionSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestMinHeapSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.MinHeapSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("MinHeap() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestMaxHeapSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.MaxHeapSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("MaxHeap() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.BubbleSort()
		b.StopTimer()
		s.data = data[:]
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.SelectionSort()
		b.StopTimer()
		s.data = data[:]
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.InsertionSort()
		b.StopTimer()
		s.data = data[:]
	}
}

func BenchmarkMinHeap(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MinHeapSort()
		b.StopTimer()
		s.data = data[:]
	}
}

func BenchmarkMaxHeap(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MaxHeapSort()
		b.StopTimer()
		s.data = data[:]
	}
}