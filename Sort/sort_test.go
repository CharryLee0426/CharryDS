package sort

import (
	"math/rand"
	"sort"
	"testing"
)

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

func TestQuickSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.QuickSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("QuickSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.MergeSort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("QuickSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

func TestPDQSort(t *testing.T) {
	inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	expected := [...]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	s := NewSort(inputs[:])

	s.PDQsort()

	for i := 0; i < s.Len(); i++ {
		if s.data[i] != expected[i] {
			t.Errorf("QuickSort() = %d; expected %d", s.data[i], expected[i])
		}
	}
}

// 36x Shuffled data
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

		// reset data to unsorted state
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
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
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
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
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
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
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
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
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.QuickSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MergeSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
	}
}

func BenchmarkPDQSort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.PDQsort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
	}
}

func BenchmarkOfficialPDQsort(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
	}


	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		sort.Ints(s.data)
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i ^ 0x2cc
		}
	}
}

// Sorted data
func BenchmarkBubbleSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.BubbleSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkSelectionSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.SelectionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkInsertionSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.InsertionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkMinHeapSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MinHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkMaxHeapSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MaxHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkQuickSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.QuickSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkMergeSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MergeSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkPDQSortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.PDQsort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

func BenchmarkPdqsortSorted(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		sort.Ints(s.data)
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i
		}
	}
}

// Reversed data
func BenchmarkBubbleSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.BubbleSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkSelectionSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.SelectionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkInsertionSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.InsertionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkMinHeapReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MinHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkMaxHeapReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MaxHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkQuickSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.QuickSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkMergeSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MergeSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkPDQSortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.PDQsort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

func BenchmarkOfficialPDQsortReversed(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = len(data) - i
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		sort.Ints(s.data)
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = len(data) - i
		}
	}
}

// Mod8 data
func BenchmarkBubbleSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.BubbleSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkSelectionSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.SelectionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkInsertionSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.InsertionSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkMinHeapMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MinHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkMaxHeapMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MaxHeapSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkQuickSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.QuickSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkMergeSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MergeSort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkPDQSortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.PDQsort()
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

func BenchmarkOfficialPDQsortMod8(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i % 8
	}

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		sort.Ints(s.data)
		b.StopTimer()
		for i := 0; i < len(data); i++ {
			s.data[i] = i % 8
		}
	}
}

// Random data
func BenchmarkBubbleSortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.BubbleSort()
		b.StopTimer()
		
		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkSelectionSortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.SelectionSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })


	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.InsertionSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkMinHeapRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })


	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MinHeapSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkMaxHeapRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })


	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MaxHeapSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkQuickSortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.QuickSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkMergeSortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.MergeSort()
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkPDQSortRandom(b *testing.B) {
	// inputs := [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
			data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s.PDQsort()
		b.StopTimer()
		
		rand.Shuffle(len(data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}

func BenchmarkOfficialPDQsortRandom(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	s := NewSort(data[:])

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		sort.Ints(s.data)
		b.StopTimer()

		rand.Shuffle(len(s.data), func(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] })
	}
}