package graph

// SPFA is an algorithm that is still based on the principle of relaxation,
// but it is more efficient than Bellman-Ford algorithm.

type SPFA struct {
	Graph [][]int
	Min []int
	N int
}


func NewSPFA(graph [][]int, n int) *SPFA {
	s := &SPFA {
		Graph: graph,
		Min: make([]int, n),
		N: n,
	}

	for i := range s.Min {
		s.Min[i] = Inf
	}

	return s
}

func (s *SPFA) Solve(source int) {
	if source < 0 || source >= s.N {
		return
	}

	s.Min[source] = 0
	queue := []int{source}

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		for _, edge := range s.Graph {
			from, to, weight := edge[0], edge[1], edge[2]

			if start == from && s.Min[to] > s.Min[from] + weight {
				s.Min[to] = s.Min[from] + weight
				queue = append(queue, to)
			}
		}
	}
}