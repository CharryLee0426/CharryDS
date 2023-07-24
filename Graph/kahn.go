package graph

// Kahn algorithm is used to do topological sorting task
// The intuition behind this khan algorithm is recursion
// At the beginning, you need to know the in-degree of
// each node. Then we choose a node whose in-degree is
// zero as the start, update the in-degree table, and
// repeat the process above until we can't do it. If
// we can find a route contains all nodes, we find a
// topological sorting route. The implementation depends
// on an in-degree array, a queue, and an array for storing
// route.

type Kahn struct {
	Graph [][]int
	InDegree []int
	N int
}

// we assumed that the graph is a directed acyclic graph
// and the graph is represented by edge list ([from, to])
// if the graph is represented by adjacency matrix, thing
// will be much easier
func NewKahn(graph [][]int, n int) *Kahn {
	inDegree := make([]int, n)

	for _, edge := range graph {
		to := edge[1]
		inDegree[to]++
	}

	return &Kahn {
		Graph: graph,
		InDegree: inDegree,
		N: n,
	}
}

// we will return an empty array if we can't find a route
func (k *Kahn) Solve() []int {
	if k.N == 1 {
		return []int{0}
	}

	queue := []int{}

	for i := range k.InDegree {
		if k.InDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	index := 0
	route := make([]int, k.N)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		route[index] = n
		index++

		for _, edge := range k.Graph {
			from := edge[0]
			to := edge[1]

			if from == n {
				k.InDegree[to]--

				if k.InDegree[to] == 0 {
					queue = append(queue, to)
				}
			}
		}
	}

	if index < k.N {
		return []int{}
	} else {
		return route
	}
}