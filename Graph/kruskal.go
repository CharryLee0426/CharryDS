package graph

import "sort"

// Kruskal's algorithm is used to find the minimum spanning tree
// of a graph. Its performance is better if there are not many edges.
// To implement it, we need a disjoint set data structure
// (also called "union find"). In order to get the minimum edge every
// time, we need to sort the edges by weight.

// Union Find data structure optimized with rank and path compression

type UnionFind struct {
	Root []int
	Rank []int
	N int
}

func NewUnionFind(n int) *UnionFind {
	root := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 0
	}

	return &UnionFind{root, rank, n}
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.Root[x] {
		x = uf.Root[x]
	}

	return x
}

func (uf *UnionFind) Union(x, y int) {
	if x >= uf.N || y >= uf.N || x < 0 || y < 0 {
		return
	}

	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}

	if uf.Rank[rootX] < uf.Rank[rootY] {
		uf.Root[rootX] = rootY
	} else if uf.Rank[rootX] > uf.Rank[rootY] {
		uf.Root[rootY] = rootX
	} else {
		uf.Root[rootY] = rootX
		uf.Rank[rootX]++
	}
}

func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Kruskal's data structure begins here
// We assume that this graph is an undirected graph
// [node1, node2, weight]
type Kruskal struct {
	Graph [][]int
	Tree [][]int
	N int
}

func NewKruskal(graph [][]int, n int) *Kruskal {
	return &Kruskal{graph, make([][]int, 0), n}
}

func (ks *Kruskal) Solve() {
	uf := NewUnionFind(ks.N)
	

	// sort the graph by weight
	sort.Slice(ks.Graph, func (i, j int) bool {
		return ks.Graph[i][2] < ks.Graph[j][2]
	})

	// there is space for improving
	for _, edge := range ks.Graph {
		if !uf.IsConnected(edge[0], edge[1]) {
			ks.Tree = append(ks.Tree, edge)
			uf.Union(edge[0], edge[1])
		}
	}
}