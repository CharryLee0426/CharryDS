package heap

import (
	"fmt"
	"math"
)

// we create a min heap whose elements are integers

type Element int

type Heap struct {
	elements []Element
	maxSize int
	realSize int
}

// Constructer
func New(maxSize int) *Heap {
	return &Heap{
		elements: make([]Element, maxSize+1),	// we don't use the first element
		maxSize: maxSize,
		realSize: 0,
	}
}

// get the size of this heap right now
func (h *Heap) Size() int {
	return h.realSize
}

// Insert an element into the heap
func (h *Heap) Insert(e Element) {
	if h.realSize >= h.maxSize {
		fmt.Println("The heap is full!")
		return
	}
	h.realSize++
	h.elements[h.realSize] = e

	// adjust the heap to make it a min heap
	i := h.realSize
	for i > 1 {
		if h.elements[i] < h.elements[i/2] {
			h.elements[i], h.elements[i/2] = h.elements[i/2], h.elements[i]
			i /= 2
		} else {
			break
		}
	}
}

// Peek the top element of the heap
func (h *Heap) Peek() Element {
	if h.realSize == 0 {
		fmt.Println("The heap is empty!")
		return Element(0)
	}
	return h.elements[1]
}

// Pop the top element of the heap
func (h *Heap) Pop() Element {
	if h.realSize == 0 {
		fmt.Println("The heap is empty!")
		return Element(0)
	}

	result := h.elements[1]

	// adjust the heap to make it a min heap
	h.elements[1] = h.elements[h.realSize]
	h.realSize--
	i := 1
	for i*2 <= h.realSize {
		if h.elements[i] > h.elements[i*2] || 
		   (h.elements[i] > h.elements[i*2+1] && i*2+1 <= h.realSize) {
			if h.elements[i*2] < h.elements[i*2+1] || i*2+1 > h.realSize {
				h.elements[i], h.elements[i*2] = h.elements[i*2], h.elements[i]
				i *= 2
			} else {
				h.elements[i], h.elements[i*2+1] = h.elements[i*2+1], h.elements[i]
				i = i*2 + 1
			}
		} else {
			break
		}
	}

	return result
}

// print the heap like a tree
func (h *Heap) Print() {
	layer := math.Floor(math.Log2(float64(h.realSize)))
	cur_layer := 0

	for i := 1; i <= h.realSize; {
		for j := 0; j < int(layer); j++ {
			fmt.Print(" ")
		}

		for k := 0; k < int(math.Min(float64(h.realSize+1)-math.Pow(2, float64(cur_layer)), math.Pow(2, float64(cur_layer)))); k++ {
			fmt.Print(h.elements[i+k], " ")
		}

		fmt.Println()

		layer--
		i = i + int(math.Pow(2, float64(cur_layer)))
		cur_layer++
	}
}