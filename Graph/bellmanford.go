package graph

// Bellman-Ford algorithm can be used to find the shortest path from a source
// vertex to all other vertices in a weighted graph.
// It is more versatile than Dijkstra's algorithm because the edges of the graph
// can have negative weights.

// The algorithm is based on the principle of relaxation, in which an estimate
// of the shortest path from the source to a vertex is gradually replaced by a
// more accurate one until eventually reaching the optimum solution.

// In fact, the implementation shows that the algorithm is based on dynamic
// programming, because it solves the problem by combining solutions to
// subproblems.

// How we store our directed weighted graph:
// we use a 2-dimensional array to store the graph.

const Inf = 1 << 30

type BellmanFord struct {
	Graph [][]int
	Dp [][]int
	Min []int
	N int
}

func NewBellmanFord(graph [][]int, n int) *BellmanFord {
	b := &BellmanFord{
		Graph: graph,
		Dp: make([][]int, n),
		Min: make([]int, n),
		N: n,
	}

	for i := range b.Dp {
		b.Dp[i] = make([]int, n)

		for j := range b.Dp[i] {
			b.Dp[i][j] = Inf
		}
	}

	for i := range b.Min {
		b.Min[i] = Inf
	}

	return b
}

func (b *BellmanFord) Solve(source int) {
	if source  < 0 || source >= b.N {
		return
	}

	b.Dp[0][source] = 0
	b.Min[source] = 0

	for i := 1; i < b.N; i++ {
		for _, edge := range b.Graph {
			from, to, weight := edge[0], edge[1], edge[2]

			b.Dp[i][to] = min(b.Dp[i][to], b.Dp[i - 1][from] + weight)
			b.Min[to] = min(b.Min[to], b.Dp[i][to])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}