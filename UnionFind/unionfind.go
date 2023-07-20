package unionfind

// Union Find data structure is very helpful when
// judging whether two nodes are connected in a graph.
// This data structure is optimized by union by rank 
// and path compression.

type UnionFind struct {
	root []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		root: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.root[i] = i
		uf.rank[i] = 1
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] == x {
		return x
	} else {
		uf.root[x] = uf.Find(uf.root[x])
		return uf.root[x]
	}
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)

	if x == y {
		return
	}
	
	if uf.rank[x] < uf.rank[y] {
		uf.root[x] = y
	} else {
		uf.root[y] = x
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}