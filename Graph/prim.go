package graph

// Prim Algorithm is used to find the minimum
// spanning tree of a graph.

type Prim struct {
	Matrix [][]int
	Visited map[int]bool // It'll be used like a set
	Tree [][]int
	N int
}

// I assume that the graph is stored in a 2D array.
// Each element stores start, end, and weight of an
// edge.

func NewPrim(graph [][]int, n int) *Prim {
	matrix := make([][]int, n)
	
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range graph[i] {
			matrix[i][j] = Inf
		}
	}

	for _, edge := range graph {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	return &Prim{matrix, make(map[int]bool, 0), [][]int{}, n}
}

func (p *Prim) Solve() int {
	// we will select the first node as start
	p.Visited[0] = true

	// the total weight of the minimum spanning tree
	total := 0

	// we will loop until we visit all the nodes
	for len(p.Visited) < p.N {
		// we will select the minimum weight edge
		// that connects a visited node and an unvisited node
		min := Inf
		from := -1
		to := -1

		for node := range p.Visited {
			for i := 0; i < p.N; i++ {
				if !p.Visited[i] && p.Matrix[node][i] < min {
					min = p.Matrix[node][i]
					from = node
					to = i
				}
			}
		}


		// we will add the weight of the edge to the total
		total += min

		p.Tree = append(p.Tree, []int{from, to, min})

		// we will mark the node as visited
		p.Visited[to] = true
	}

	return total
}