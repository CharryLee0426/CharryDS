# Charry Data Structure

## 0. Overview
Charry Data Structure is a repository that stores some useful data structures.
I use go programming language to implement these data structures due to three
advantages it has. First, go is a compiled language, which means it is faster
than many interpreted languages such as python. Second, it is very easy for
developers to design tests for their code. I love go test! Third, go is a
very widely used language right now, especially when it comes to cloud.

## 1. Data Structures
* Heap (min heap)
* Binary Tree (assumed to create a complete binary tree)

## 2. Algorithms
* Sort
    - [x] Heap Sort
    - [x] Insertion Sort
    - [x] Selection Sort
    - [x] Bubble Sort
    - [ ] Quick Sort
    - [ ] Merge Sort

Test condition:
Three types of data: shuffle, sorted, reversed, mod 8 (1k elements)
Disable compiler optimization
100 times for each type of data
Show memory allocation and number of allocations

Test result:
```
goos: darwin
goarch: arm64
pkg: charryds/Sort
BenchmarkBubbleSort-8                       1000           2484494 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-8                    1000           1232665 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-8                    1000           2047467 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeap-8                          1000           1025436 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeap-8                          1000            101511 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-8                        1000             55108 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsort-8                          1000             33365 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortSorted-8                 1000           1169204 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortSorted-8              1000           1172579 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortSorted-8              1000              3156 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapSorted-8                    1000            111541 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapSorted-8                    1000            106992 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortSorted-8                  1000           1143183 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortSorted-8                    1000              2655 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortReversed-8               1000           2976828 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortReversed-8            1000           1196626 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortReversed-8            1000           2943280 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapReversed-8                  1000           1031044 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapReversed-8                  1000            100347 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortReversed-8                1000           1317937 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortReversed-8                  1000              3889 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortMod8-8                   1000           1214647 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortMod8-8                1000           1189833 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortMod8-8                1000           1279064 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapMod8-8                      1000            714547 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapMod8-8                      1000             95185 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortMod8-8                    1000            193249 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortMod8-8                      1000             13559 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortRandom-8                 1000           2642066 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortRandom-8              1000           1227392 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortRandom-8              1000           1468509 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapRandom-8                    1000            748422 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapRandom-8                    1000            114664 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortRandom-8                  1000             73856 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortRandom-8                    1000             63525 ns/op              24 B/op          1 allocs/op
```
