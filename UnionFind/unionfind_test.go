package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	expected := []bool{true, true, false, true}

	uf := NewUnionFind(10)

	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)

	actual := []bool{
		uf.IsConnected(1, 5),
		uf.IsConnected(5, 7),
		uf.IsConnected(4, 9),
	}

	uf.Union(9, 4)

	actual = append(actual, uf.IsConnected(4, 9))

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected[i], actual[i])
		}
	}
}