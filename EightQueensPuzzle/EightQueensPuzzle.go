package eightqueenspuzzle

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

var results = make([][]Point, 0)

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
	fmt.Print("Results:\n")
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("There were %d results\n", len(results))
}


func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)

		if len(current) == n {
			c := make([]Point, n)
			for i, p := range current {
				c[i] = p
			}
			results = append(results, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)
				}
			}
		}
	}
}

// CanPlace returns true if the target point can be placed on the board.
// Target is what we want to place, board is the set of queens we have placed.
// We check each point in the board to see if it can attack the target point.
// If so, we just escape the loop and refure false. Otherwise, return true.
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

// CanAttack returns true if the two points are in the same row, column, 
// or diagonal. It's very smart that using the absolute value to check if
// the two points are in the same diagonal.
func CanAttack(a, b Point) bool {
	attack := a.x == b.x || a.y == b.y || math.Abs(float64(a.x-b.x)) == math.Abs(float64(a.y-b.y))
	return attack
}