package graph

import "math"

/*
Floyd-Warshall algorithm is an algorithm for finding shortest paths in a weighted graph
with positive or negative edge weights (but with no negative cycles).
Floyd-Warshall algorithm is a good example of dynamic programming.
*/

type FloydWarshall struct {
	Graph [][]int
	N int
}

func NewFloydWarshall(edges [][]int, n int) FloydWarshall {
	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = math.MaxInt32 / 2
		}
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from][to], graph[to][from] = weight, weight
	}

	return FloydWarshall {
		Graph: graph,
		N: n,
	}
}

func (fw *FloydWarshall) Solve() [][]int {
	n := fw.N

	for k := 0; k < n; k++ {
		fw.Graph[k][k] = 0

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fw.Graph[i][j] = min(fw.Graph[i][j], fw.Graph[i][k]+fw.Graph[k][j])
			}
		}
	}

	return fw.Graph
}