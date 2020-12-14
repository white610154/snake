package main

import "math/rand"

// Point 2D position or vector
type Point struct {
	X int
	Y int
}

// RandomPoint generate a random point
func RandomPoint(avoid ...Point) Point {
	p := Point{rand.Intn(fieldWidth) + 1, rand.Intn(fieldHeight) + 1}
	for {
		crash := false
		for _, q := range avoid {
			if IsEqual(p, q) {
				crash = true
				break
			}
		}
		if crash {
			p.X = rand.Intn(fieldWidth) + 1
			p.Y = rand.Intn(fieldHeight) + 1
		} else {
			break
		}
	}
	return p
}

// IsEqual return true if 2 points are the same
func IsEqual(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}
