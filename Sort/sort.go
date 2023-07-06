package sort

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

// TODO: pdqsort