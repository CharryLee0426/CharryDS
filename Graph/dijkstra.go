package graph

import "container/heap"

// First we need to implement a priority queue. We'll use a heap

type Node struct {
	Distance int
	Index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(a interface{}) {
	*pq = append(*pq, a.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	a := old[n-1]
	*pq = old[0 : n-1]
	return a
}

// Now we can implement Dijkstra's algorithm
// The intuition is the greddy algorithm: always choose the shortest path

type Dijkstra struct {
	Net map[int][]Node
	Min map[int]int
	N int
}

func NewDijkstra(graph [][]int, n int) *Dijkstra {
	net := make(map[int][]Node)

	for _, edge := range graph {
		u, v, w := edge[0], edge[1], edge[2]
		net[u] = append(net[u], Node{w, v})
	}

	min := make(map[int]int)

	return &Dijkstra{net, min, n}
}

func (d *Dijkstra) Solve(source int) {
	d.Min[source] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, &Node{0, source})

	for pq.Len() > 0 {
		n := heap.Pop(pq).(*Node)

		if _, ok := d.Min[n.Index]; ok && d.Min[n.Index] < n.Distance {
			continue
		}

		d.Min[n.Index] = n.Distance
		
		for _, target := range d.Net[n.Index] {
			targetIndex := target.Index
			targetDis := target.Distance + n.Distance

			heap.Push(pq, &Node{targetDis, targetIndex})
		}
	}
}