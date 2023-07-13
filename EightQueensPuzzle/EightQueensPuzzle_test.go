package eightqueenspuzzle

import "testing"

func BenchmarkEightQueensPuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(8)
	}
}