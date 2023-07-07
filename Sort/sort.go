package sort

import "math/bits"

// Assume that all elements are integers.
// Assume we need to sort in ascending order.

type Sort struct {
	data []int
	len int
	sortedData []int
}

func NewSort(data []int) *Sort {
	return &Sort{data: data, len: len(data)}
}

func (s *Sort) Len() int {
	return s.len
}

func (s *Sort) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

func (s *Sort) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *Sort) IsSorted() bool {
	if s.sortedData == nil {
		return false
	}

	return true
}

func (s *Sort) Sorted() []int {
	return s.sortedData
}

// Sorting Algorithms

// Bubble sort is a simple sorting algorithm
// Time complexity: n-1 + n-2 + ... + 1 =
// n(n-1)/2 ===> O(n^2)
// Space complexity: O(1)
// Stable: Yes
func (s *Sort) BubbleSort() {
	for i := 0; i < s.len; i++ {
		for j := i + 1; j < s.len; j++ {
			if s.data[i] > s.data[j] {
				s.Swap(i, j)
			}
		}
	}

	s.sortedData = s.data
}

// Selection sort is a simple sorting algorithm
// Time complexity: n-1 + n-2 + ... + 1 =
// n(n-1)/2 ===> O(n^2)
// Space complexity: O(1)
// Stable: No
func (s *Sort) SelectionSort() {
	for i := 0; i < s.len; i++ {
		min := i
		for j := i + 1; j < s.len; j++ {
			if s.data[j] < s.data[min] {
				min = j
			}
		}

		s.Swap(i, min)
	}

	s.sortedData = s.data
}

// Insertion sort is a simple sorting algorithm
// Time complexity: best case: O(n), worst case: O(n^2)
// Space complexity: O(1)
// Stable: Yes
func (s *Sort) InsertionSort() {
	for i := 1; i < s.len; i++ {
		for j := i; j > 0 && s.data[j] < s.data[j-1]; j-- {
			s.Swap(j, j-1)
		}
	}

	s.sortedData = s.data
}

// Heap sort is a simple sorting algorithm
// Time complexity: O(nlogn)
// Space complexity: O(n)
// Stable: No
func (s *Sort) MinHeapSort() {
	h := newHeap(s.data, s.len)

	for i := 0; i < s.len; i++ {
		s.data[i] = h.Pop()
	}

	s.sortedData = s.data
}

// tool function used in heap sort
type heap struct {
	elements []int
	maxSize int
	realSize int
}

func newHeap(data []int, maxSize int) *heap {
	if len(data) > maxSize {
		return nil
	}

	h := &heap{
		elements: make([]int, maxSize+1),
		maxSize: maxSize,
		realSize: len(data),
	}

	for i := 1; i <= len(data); i++ {
		h.elements[i] = data[i-1]
	}

	heaptify(h.elements, h.realSize)

	return h
}

func heaptify(elements []int, realSize int) {
	for i := realSize/2; i >= 1; i-- {
		// it assumed that it's not a leaf node
		for i*2 <= realSize {
			if elements[i] > elements[i*2] || 
			   (i*2+1 <= realSize && elements[i] > elements[i*2+1]) {
				if i*2+1 > realSize || elements[i*2] < elements[i*2+1] {
					elements[i], elements[i*2] = elements[i*2], elements[i]
					i *= 2
				} else {
					elements[i], elements[i*2+1] = elements[i*2+1], elements[i]
					i = i*2 + 1
				}
			} else {
				break
			}
		}
	}
}

func (h *heap) Pop() int {
	if h.realSize == 0 {
		return h.elements[0]
	}

	result := h.elements[1]

	// adjust the heap to make it a min heap
	h.elements[1] = h.elements[h.realSize]
	h.realSize--
	i := 1
	for i*2 <= h.realSize {
		if h.elements[i] > h.elements[i*2] || 
		   (i*2+1 <= h.realSize && h.elements[i] > h.elements[i*2+1]) {
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

// another implementation of heap sort inspired by go
func slipDown(data []int, lo, hi, first int) {
	root := lo

	for {
		child := root * 2 + 1
		
		if child >= hi {
			break
		}

		if child + 1 < hi && data[first+child] < data[first+child+1] {
			child++
		}

		if data[first+root] >= data[first+child] {
			return
		}
		
		// swap
		data[first+root], data[first+child] = data[first+child], data[first+root]

		root = child
	}
}

func (s *Sort) MaxHeapSort() {
	// heaptify the data as a max heap
	for i := s.len/2 - 1; i >= 0; i-- {
		slipDown(s.data, i, s.len, 0)
	}

	for i := s.len - 1; i >= 0; i-- {
		s.data[0], s.data[i] = s.data[i], s.data[0]
		slipDown(s.data, 0, i, 0)
	}

	s.sortedData = s.data
}

func (s *Sort) QuickSort() {
	quickSort(s.data, 0, s.len-1)
	s.sortedData = s.data
}

func quickSort(nums []int, beginIndex, endIndex int) {
    if beginIndex < endIndex {
        pivot := nums[endIndex]
        i, j := beginIndex, endIndex
        for i != j {
            if nums[i] <= pivot {
                i++
            } else if nums[j] >= pivot {
                j--
            } else {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }

        // let pivot to its right place
        nums[j], nums[endIndex] = nums[endIndex], nums[j]

        // recurse
        quickSort(nums, beginIndex, i-1)
        quickSort(nums, i+1, endIndex)
    } else {
        return
    }
}

func (s *Sort) MergeSort() {
	mergeSort(s.data, 0, s.len-1)
	s.sortedData = s.data
}

func mergeSort(nums []int, beginIndex, endIndex int) {
	if beginIndex < endIndex {
		mid := (beginIndex + endIndex) / 2
		mergeSort(nums, beginIndex, mid)
		mergeSort(nums, mid+1, endIndex)
		merge(nums, beginIndex, mid, endIndex)
	}
}

// merge two sorted array to one big sorted array
func merge(nums []int, beginIndex, mid, endIndex int) {
	tmp := make([]int, endIndex-beginIndex+1)
	i, j, k := beginIndex, mid+1, 0

	for i <= mid && j <= endIndex {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= endIndex {
		tmp[k] = nums[j]
		j++
		k++
	}

	// we use a temp array to store
	// merged array temporarily
	for i := 0; i < len(tmp); i++ {
		nums[beginIndex+i] = tmp[i]
	}
}

func (s *Sort) PDQsort() {
	n := s.Len()
	if n <= 1 {
		return
	}

	limit := bits.Len(uint(n))
	pdqsort(s.data, 0, n, limit)
	s.sortedData = s.data
}

func pdqsort(data []int, a, b, limit int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true // whether the last partitioning was reasonably balanced
		wasPartitioned = true // whether the slice was already partitioned
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSort(data, a, b)
			return
		}

		// Fall back to heapsort if too many bad choices were made.
		if limit == 0 {
			heapSort(data, a, b)
			return
		}

		// If the last partitioning was imbalanced, we need to breaking patterns.
		if !wasBalanced {
			breakPatterns(data, a, b)
			limit--
		}

		pivot, hint := choosePivot(data, a, b)
		if hint == decreasingHint {
			reverseRange(data, a, b)
			
			pivot = (b - 1) - (pivot - a)
			hint = increasingHint
		}

		// The slice is likely already sorted.
		if wasBalanced && wasPartitioned && hint == increasingHint {
			if partialInsertionSort(data, a, b) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if a > 0 && !(data[a-1] < data[pivot]) {
			mid := partitionEqual(data, a, b, pivot)
			a = mid
			continue
		}

		mid, alreadyPartitioned := partition(data, a, b, pivot)
		wasPartitioned = alreadyPartitioned

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8
		if leftLen < rightLen {
			wasBalanced = leftLen >= balanceThreshold
			pdqsort(data, a, mid, limit)
			a = mid + 1
		} else {
			wasBalanced = rightLen >= balanceThreshold
			pdqsort(data, mid+1, b, limit)
			b = mid
		}
	}
}

const (
	unknownHint = iota
	increasingHint
	decreasingHint
)

// tool functions used in pdqsort
func insertionSort(data []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func heapSort(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// build a max heap
	for i := (hi - 1) / 2; i >= 0; i-- {
		slipDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		data[first], data[first+i] = data[first+i], data[first]
		slipDown(data, lo, i, first)
	}
}

type xorshift uint64

func (r *xorshift) Next() uint64 {
	*r ^= *r << 13
	*r ^= *r >> 17
	*r ^= *r << 5
	return uint64(*r)
}

func nextPowerOfTwo(length int) uint {
	shift := uint(bits.Len(uint(length)))
	return uint(1 << shift)
}

func breakPatterns(data []int, a, b int) {
	length := b - a
	if length >= 8 {
		random := xorshift(length)
		modulus := nextPowerOfTwo(length)

		for idx := a + (length/4)*2 - 1; idx <= a+(length/4)*2+1; idx++ {
			other := int(uint(random.Next()) & (modulus - 1))
			if other >= length {
				other -= length
			}
			data[idx], data[a+other] = data[a+other], data[idx]
		}
	}
}

func choosePivot(data []int, a, b int) (pivot int, hint int) {
	const (
		shortestNinther = 50
		maxSwaps        = 4 * 3
	)

	l := b - a

	var (
		swaps int
		i     = a + l/4*1
		j     = a + l/4*2
		k     = a + l/4*3
	)

	if l >= 8 {
		if l >= shortestNinther {

			i = medianAdjacent(data, i, &swaps)
			j = medianAdjacent(data, j, &swaps)
			k = medianAdjacent(data, k, &swaps)
		}

		j = median(data, i, j, k, &swaps)
	}

	switch swaps {
	case 0:
		return j, increasingHint
	case maxSwaps:
		return j, decreasingHint
	default:
		return j, unknownHint
	}
}

func order2(data []int, a, b int, swaps *int) (int, int) {
	if data[b] < data[a] {
		*swaps++
		return b, a
	}

	return a, b
}

func median(data []int, a, b, c int, swaps *int) int {
	a, b = order2(data, a, b, swaps)
	b, c = order2(data, b, c, swaps)
	a, b = order2(data, a, b, swaps)
	return b
}

func medianAdjacent(data []int, a int, swaps *int) int {
	return median(data, a-1, a, a+1, swaps)
}

func reverseRange(data []int, a, b int) {
	i := a
	j := b - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func partition(data []int, a, b, pivot int) (newpivot int, alreadyPartitioned bool) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a+1, b-1 // i and j are inclusive of the elements remaining to be partitioned

	for i <= j && data[i] < data[a] {
		i++
	}
	for i <= j && !(data[j] < data[a]) {
		j--
	}
	if i > j {
		data[j], data[a] = data[a], data[j]
		return j, true
	}
	data[i], data[j] = data[j], data[i]
	i++
	j--

	for {
		for i <= j && data[i] < data[a] {
			i++
		}
		for i <= j && !(data[j] < data[a]) {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	data[j], data[a] = data[a], data[j]
	return j, false
}

func partitionEqual(data []int, a, b, pivot int) (newpivot int) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a+1, b-1

	for {
		for i <= j && !(data[a] < data[i]) {
			i++
		}
		for i <= j && data[a] < data[j] {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return i
}

func partialInsertionSort(data []int, a, b int) bool {
	const (
		maxSteps         = 5  // maximum number of adjacent out-of-order pairs that will get shifted
		shortestShifting = 50 // don't shift any elements on short arrays
	)
	i := a + 1
	for j := 0; j < maxSteps; j++ {
		for i < b && !(data[i] < data[i-1]) {
			i++
		}

		if i == b {
			return true
		}

		if b-a < shortestShifting {
			return false
		}

		data[i], data[i-1] = data[i-1], data[i]

		// Shift the smaller one to the left.
		if i-a >= 2 {
			for j := i - 1; j >= 1; j-- {
				if !(data[j] < data[j-1]) {
					break
				}
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		// Shift the greater one to the right.
		if b-i >= 2 {
			for j := i + 1; j < b; j++ {
				if !(data[j] < data[j-1]) {
					break
				}
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return false
}