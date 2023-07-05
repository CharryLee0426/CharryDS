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

Test result:
```
goos: darwin
goarch: arm64
pkg: charryds/Sort
BenchmarkBubbleSort-8               3358            357092 ns/op
BenchmarkSelectionSort-8            2851            418822 ns/op
BenchmarkInsertionSort-8         1459506               821.6 ns/op
BenchmarkMinHeap-8                 50012             22182 ns/op
BenchmarkMaxHeap-8                 82112             14140 ns/op
PASS
ok      charryds/Sort   106.796s
```
