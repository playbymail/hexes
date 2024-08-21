// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package hexes

// Hex implements an immutable, regular hexagon using "cube coordinates."
type Hex struct {
	q, r, s int
}

// NewHex returns a new hex with the given coordinates.
func NewHex(q, r, s int) Hex {
	if q+r+s != 0 {
		panic("assert(q + r + s == 0)")
	}
	return Hex{q: q, r: r, s: s}
}

// FractionalHex implements an immutable, regular hexagon using "cube coordinates."
type FractionalHex struct {
	q, r, s float64
}

// --------------------------------------------------------------------------------------------------------------------
// equality

// Equals returns true if the two hexes have the same coordinates.
func Equals(a, b Hex) bool {
	return a.q == b.q && a.s == b.s && a.r == b.r
}

// --------------------------------------------------------------------------------------------------------------------
// arithmetic

func hex_add(a, b Hex) Hex {
	return Hex{q: a.q + b.q, r: a.r + b.r, s: a.s + b.s}
}

func hex_subtract(a, b Hex) Hex {
	return Hex{q: a.q - b.q, r: a.r - b.r, s: a.s - b.s}
}

func hex_multiply(a Hex, k int) Hex {
	return Hex{q: a.q * k, r: a.r * k, s: a.s * k}
}

// --------------------------------------------------------------------------------------------------------------------
// distance

// abs returns the absolute value of i.
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func hex_length(hex Hex) int {
	return (abs(hex.q) + abs(hex.r) + abs(hex.s)) / 2
}

func hex_distance(a, b Hex) int {
	return hex_length(hex_subtract(a, b))
}

// --------------------------------------------------------------------------------------------------------------------
// neighbors

// hex_directions is counter-clockwise starting from the "southeast."
var hex_directions = [6]Hex{
	{q: 1, r: 0, s: -1}, //   0 -> southeast    30 degrees
	{q: 1, r: -1, s: 0}, //   1 -> northeast   330 degrees
	{q: 0, r: -1, s: 1}, //   2 -> north       270 degrees
	{q: -1, r: 0, s: 1}, //   3 -> northwest   210 degrees
	{q: -1, r: 1, s: 0}, //   4 -> southwest   150 degrees
	{q: 0, r: 1, s: -1}, //   5 -> south        90 degrees
}

func hex_direction(direction int /* 0 to 5 */) Hex {
	direction = (6 + (direction % 6)) % 6
	return hex_directions[direction]
}

func hex_neighbor(hex Hex, direction int) Hex {
	return hex_add(hex, hex_direction(direction))
}

func (l *Layout_t) OffsetToHex(col, row int) Hex {
	r := row - (col+(col&1))/2
	return Hex{q: col, r: r, s: -col - r}
}

type OffsetHex_t struct {
	col, row int
}

// --------------------------------------------------------------------------------------------------------------------

// FlatHex_t implements an immutable, regular FlatTop hexagon using "cube coordinates."
type FlatHex_t struct {
	Q, R, S int
}

// NewFlatHex returns a new hex with the given coordinates.
func NewFlatHex(q, r, s int) FlatHex_t {
	if q+r+s != 0 {
		panic("assert(q + r + s == 0)")
	}
	return FlatHex_t{Q: q, R: r, S: s}
}

func (h FlatHex_t) Add(b FlatHex_t) FlatHex_t {
	return NewFlatHex(h.Q+b.Q, h.R+b.R, h.S+b.S)
}

// DiagonalNeighbor returns the neighbor in the given direction.
// The direction is an integer from 0 to 5, with "north" being 0 and "northwest" being 5.
// The direction is allowed to wrap left or right, so values outside the range [0, 5] are allowed.
func (h FlatHex_t) DiagonalNeighbor(direction int) FlatHex_t {
	return h.Add(flatHexDiagonals[(6+(direction%6))%6])
}

// flatHexDiagonals is the list of offsets for the six diagonals of a FlatTop hexagon.
// The list is in the order of the directions, starting with "north" and going clockwise.
var flatHexDiagonals = []FlatHex_t{
	NewFlatHex(2, -1, -1), // north
	NewFlatHex(1, -2, 1),  // northeast
	NewFlatHex(-1, -1, 2), // southeast
	NewFlatHex(-2, 1, 1),  // south
	NewFlatHex(-1, 2, -1), // southwest
	NewFlatHex(1, 1, -2),  // northwest
}

func (h FlatHex_t) RotateLeft() FlatHex_t {
	return NewFlatHex(-h.S, -h.Q, -h.R)
}

func (h FlatHex_t) RotateRight() FlatHex_t {
	return NewFlatHex(-h.R, -h.S, -h.Q)
}
