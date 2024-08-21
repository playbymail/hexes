// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package points

import "math"

// Points are immutable x, y coordinates on a 2D plane.
// The origin is at the bottom left. The X-axis increases moving right and the Y-axis increases moving up.

// Point is an immutable 2D coordinate on a plane.
type Point struct {
	X, Y float64
}

// NewPoint returns a new point.
func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// Delta returns the difference between two points.
func (p Point) Delta(p2 Point) Vector2 {
	return Vector2{X: p.X - p2.X, Y: p.Y - p2.Y}
}

// Distance returns the distance between two points.
func (p Point) Distance(p2 Point) float64 {
	dx, dy := p.X-p2.X, p.Y-p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Translate returns a new point translated by the given vector.
func (p Point) Translate(v Vector2) Point {
	return Point{X: p.X + v.X, Y: p.Y + v.Y}
}

// Vector2 is an immutable 2D vector.
type Vector2 struct {
	X, Y float64
}
