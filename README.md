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
* Disjoint Set Data Structure (Union Find)

## 2. Algorithms
* Sort
    - [x] Heap Sort
    - [x] Insertion Sort
    - [x] Selection Sort
    - [x] Bubble Sort
    - [x] Quick Sort
    - [x] Merge Sort
    - [x] Pattern-defeating Quicksort (Pdqsort) (state of the art)

* Graph
    - [] Kruskal
    - [] Prim
    - [] Dijkstra
    - [x] Bellman-Ford
    - [x] SPFA

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
BenchmarkBubbleSort-8                       1000           2467353 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-8                    1000           1227912 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-8                    1000           2048261 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeap-8                          1000           1024387 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeap-8                          1000            102687 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-8                        1000             55085 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-8                        1000             82767 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSort-8                          1000             47935 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsort-8                  1000             33705 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortSorted-8                 1000           1166343 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortSorted-8              1000           1177372 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortSorted-8              1000              3175 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapSorted-8                    1000            111988 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapSorted-8                    1000            107079 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortSorted-8                  1000           1146282 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortSorted-8                  1000             83550 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortSorted-8                    1000              2652 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortSorted-8                    1000              2603 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortReversed-8               1000           2980081 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortReversed-8            1000           1199145 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortReversed-8            1000           2914308 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapReversed-8                  1000           1032363 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapReversed-8                  1000             99003 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortReversed-8                1000           1318531 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortReversed-8                1000             82179 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortReversed-8                  1000              4139 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortReversed-8          1000              3852 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortMod8-8                   1000           1210913 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortMod8-8                1000           1189246 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortMod8-8                1000           1281706 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapMod8-8                      1000            707426 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapMod8-8                      1000             96785 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortMod8-8                    1000            194367 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortMod8-8                    1000             88986 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortMod8-8                      1000             18608 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortMod8-8              1000             13765 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortRandom-8                 1000           2683426 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortRandom-8              1000           1227405 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortRandom-8              1000           1468105 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapRandom-8                    1000            748061 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapRandom-8                    1000            115578 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortRandom-8                  1000             73461 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortRandom-8                  1000            118175 ns/op           81922 B/op       1023 allocs/op
BenchmarkPDQSortRandom-8                    1000             82756 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortRandom-8            1000             63176 ns/op              24 B/op          1 allocs/op
```

* Recursive solution for eight queens puzzle