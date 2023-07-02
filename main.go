package main

import (
	heap "charryds/Heap"
	"fmt"
)



func main() {
	inputs := []int{7, 6, 5, 4, 3, 2, 1}

	h := heap.New(7)

	for _, input := range inputs {
		h.Insert(heap.Element(input))
	}

	for i := 1; i <= 7; i++ {
		h.Print()
		min := h.Pop()
		fmt.Printf("The min top is %d\n", min)
	}
}