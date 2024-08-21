// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package hexes

import (
	"fmt"
	"math"
)

// --------------------------------------------------------------------------------------------------------------------
// hexes

type Hex struct {
	q, r, s int
}

func NewHex(q, r, s int) Hex {
	return Hex{q: q, r: r, s: s}
}

// --------------------------------------------------------------------------------------------------------------------
// equality

func equals(a, b Hex) bool {
	return a.q == b.q && a.s == b.s && a.r == b.r
}

func notEquals(a, b Hex) bool {
	return !equals(a, b)
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
	return hex_directions[(6+(direction%6))%6]
}

func hex_neighbor(hex Hex, direction int) Hex {
	return hex_add(hex, hex_direction(direction))
}

// --------------------------------------------------------------------------------------------------------------------
// layout

// Orientation is a helper struct for converting between hex coordinates and screen coordinates.

type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	start_angle    float64 // in multiples of 60°
}

const (
	sqrt3 = 1.732050807568877293527446341505872366
)

var (
	layout_flat = Orientation{
		f0:          3.0 / 2.0,
		f1:          0.0,
		f2:          sqrt3 / 2.0,
		f3:          sqrt3,
		b0:          2.0 / 3.0,
		b1:          0.0,
		b2:          -1.0 / 3.0,
		b3:          sqrt3 / 3.0,
		start_angle: 0.0,
	}
	layout_pointy = Orientation{
		f0:          sqrt3,
		f1:          sqrt3 / 2.0,
		f2:          0.0,
		f3:          3.0 / 2.0,
		b0:          sqrt3 / 3.0,
		b1:          -1.0 / 3.0,
		b2:          0.0,
		b3:          2.0 / 3.0,
		start_angle: 0.5,
	}
)

type Layout interface {
	// screen conversions
	hex_to_pixel(h Hex) Point
	pixel_to_hex(p Point) FractionalHex

	// drawing
	hex_corner_offset(corner int) Point
	polygon_corners(h Hex) [6]Point

	HexToCenterPoint(h Hex) Point
	Points(h Hex) (Point, [6]Point)
}

type FlatEvenLayout struct {
	size   Point
	origin Point
}

func NewFlatEvenLayout(size, origin Point) Layout {
	return &FlatEvenLayout{
		size:   size,
		origin: origin,
	}
}

type PointyEvenLayout struct {
	size   Point
	origin Point
}

func NewPointyOddLayout(size, origin Point) Layout {
	return &PointyEvenLayout{
		size:   size,
		origin: origin,
	}
}

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

type OffsetCoord struct {
	Column, Row int
}

// --------------------------------------------------------------------------------------------------------------------
// hex to screen

func (layout *FlatEvenLayout) hex_to_pixel(h Hex) Point {
	var x = (layout_flat.f0*float64(h.q) + layout_flat.f1*float64(h.r)) * layout.size.X
	var y = (layout_flat.f2*float64(h.q) + layout_flat.f3*float64(h.r)) * layout.size.Y
	return Point{
		X: x + layout.origin.X,
		Y: y + layout.origin.Y,
	}
}

func (layout *FlatEvenLayout) HexToCenterPoint(h Hex) Point {
	return layout.hex_to_pixel(h)
}

func (layout *PointyEvenLayout) hex_to_pixel(h Hex) Point {
	var x = (layout_pointy.f0*float64(h.q) + layout_pointy.f1*float64(h.r)) * layout.size.X
	var y = (layout_pointy.f2*float64(h.q) + layout_pointy.f3*float64(h.r)) * layout.size.Y
	return Point{
		X: x + layout.origin.X,
		Y: y + layout.origin.Y,
	}
}

func (layout *PointyEvenLayout) HexToCenterPoint(h Hex) Point {
	return layout.hex_to_pixel(h)
}

// --------------------------------------------------------------------------------------------------------------------
// screen to hex

func (layout *FlatEvenLayout) pixel_to_hex(p Point) FractionalHex {
	var M = layout_flat
	var pt = Point{X: (p.X - layout.origin.X) / layout.size.X, Y: (p.X - layout.origin.Y) / layout.size.Y}
	var q = M.b0*pt.X + M.b1*pt.Y
	var r = M.b2*pt.X + M.b3*pt.Y
	return FractionalHex{q: q, r: r, s: -q - r}
}

func (layout *PointyEvenLayout) pixel_to_hex(p Point) FractionalHex {
	var M = layout_pointy
	var pt = Point{X: (p.X - layout.origin.X) / layout.size.X, Y: (p.Y - layout.origin.Y) / layout.size.Y}
	var q = M.b0*pt.X + M.b1*pt.Y
	var r = M.b2*pt.X + M.b3*pt.Y
	return FractionalHex{q: q, r: r, s: -q - r}
}

// --------------------------------------------------------------------------------------------------------------------
// drawing hex on screen

func (layout *FlatEvenLayout) hex_corner_offset(corner int) Point {
	size := layout.size
	angle := 2.0 * math.Pi * (layout_flat.start_angle + float64(corner)) / 6
	return Point{X: size.X * math.Cos(angle), Y: size.Y * math.Sin(angle)}
}

func (layout *PointyEvenLayout) hex_corner_offset(corner int) Point {
	size := layout.size
	angle := 2.0 * math.Pi * (layout_pointy.start_angle + float64(corner)) / 6
	return Point{X: size.X * math.Cos(angle), Y: size.Y * math.Sin(angle)}
}

func (layout *FlatEvenLayout) polygon_corners(h Hex) [6]Point {
	var corners [6]Point
	center := layout.hex_to_pixel(h)
	for i := 0; i < 6; i++ {
		offset := layout.hex_corner_offset(i)
		corners[i] = Point{X: center.X + offset.X, Y: center.Y + offset.Y}
	}
	return corners
}

func (layout *PointyEvenLayout) polygon_corners(h Hex) [6]Point {
	var corners [6]Point
	center := layout.hex_to_pixel(h)
	for i := 0; i < 6; i++ {
		offset := layout.hex_corner_offset(i)
		corners[i] = Point{X: center.X + offset.X, Y: center.Y + offset.Y}
	}
	return corners
}

func (layout *FlatEvenLayout) Points(h Hex) (center Point, corners [6]Point) {
	center = layout.hex_to_pixel(h)
	for i := 0; i < 6; i++ {
		offset := layout.hex_corner_offset(i)
		corners[i] = Point{X: center.X + offset.X, Y: center.Y + offset.Y}
	}
	return center, corners
}

func (layout *PointyEvenLayout) Points(h Hex) (center Point, corners [6]Point) {
	center = layout.hex_to_pixel(h)
	for i := 0; i < 6; i++ {
		offset := layout.hex_corner_offset(i)
		corners[i] = Point{X: center.X + offset.X, Y: center.Y + offset.Y}
	}
	return center, corners
}

// --------------------------------------------------------------------------------------------------------------------
// fractional hex

type FractionalHex struct {
	q, r, s float64
}

// --------------------------------------------------------------------------------------------------------------------
// rounding

func hex_round(h FractionalHex) Hex {
	q, r, s := math.Round(h.q), math.Round(h.r), math.Round(h.s)
	q_diff, r_diff, s_diff := math.Abs(q-h.q), math.Abs(r-h.r), math.Abs(s-h.s)
	if q_diff > r_diff && q_diff > s_diff {
		q = -r - s
	} else if r_diff > s_diff {
		r = -q - s
	} else {
		s = -q - r
	}
	return Hex{q: int(q), r: int(r), s: int(s)}
}

// --------------------------------------------------------------------------------------------------------------------
// line drawing

func lerp(a, b, t float64) float64 {
	//  better for floating point precision than a + (b - a) * t, which is what I usually write
	return a*(1-t) + b*t
}

func hex_lerp(a, b Hex, t float64) FractionalHex {
	q := lerp(float64(a.q), float64(b.q), t)
	r := lerp(float64(a.r), float64(b.r), t)
	s := lerp(float64(a.s), float64(b.s), t)
	return FractionalHex{q: q, r: r, s: s}
}

func hex_linedraw(a, b Hex) (results []Hex) {
	N := hex_distance(a, b)
	step := 1.0 / float64(max(N, 1))
	for i := 0; i <= N; i++ {
		results = append(results, hex_round(hex_lerp(a, b, float64(i)*step)))
	}
	return results
}

func fractional_hex_lerp(a, b FractionalHex, t float64) FractionalHex {
	q := lerp(float64(a.q), float64(b.q), t)
	r := lerp(float64(a.r), float64(b.r), t)
	s := lerp(float64(a.s), float64(b.s), t)
	return FractionalHex{q: q, r: r, s: s}
}

func hex_linedraw_with_nudge(a, b Hex) (results []Hex) {
	N := hex_distance(a, b)

	// nudge both hex centers to avoid points on the edge of the line
	afh := FractionalHex{float64(a.q) + 1e-6, float64(a.r) + 1e-6, float64(a.s) - 2e-6}
	bfh := FractionalHex{float64(b.q) + 1e-6, float64(b.r) + 1e-6, float64(b.s) - 2e-6}

	step := 1.0 / float64(max(N, 1))
	for i := 0; i <= N; i++ {
		results = append(results, hex_round(fractional_hex_lerp(afh, bfh, float64(i)*step)))
	}

	return results
}
