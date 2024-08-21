// Copyright (c) 2024 Michael D Henderson. All rights reserved.

// Package hexes implements a hexagon grid for flat hexes with odd-numbered columns pushed down.
//
// Note: the RedBlog pages call this a "flat-top odd-q" layout.
package hexes

import (
	"github.com/playbymail/hexes/points"
	"math"
)

// --------------------------------------------------------------------------------------------------------------------
// implementation based on https://www.redblobgames.com/grids/hexagons/implementation.html

const (
	sqrt3 = 1.732050807568877293527446341505872366
)

// A regular hexagon can be inscribed within two circles.
// The first is the outer circle. It passes through all six corners of the hexagon.
// The second is the inner circle. It passes through the centers of all six sides of the hexagon.

// --------------------------------------------------------------------------------------------------------------------
// layout

// Orientation is a helper struct for converting between hex coordinates and screen coordinates.

type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	start_angle    float64 // in multiples of 60°
}

var (
	layout_pointy = Orientation{
		f0:          sqrt3,
		f1:          sqrt3 / 2,
		f2:          0,
		f3:          3 / 2,
		b0:          sqrt3 / 3,
		b1:          -1 / 3,
		b2:          0,
		b3:          2 / 3,
		start_angle: 0.5,
	}
	layout_flat = Orientation{
		f0:          3 / 2,
		f1:          0,
		f2:          sqrt3 / 2,
		f3:          sqrt3,
		b0:          2 / 3,
		b1:          0,
		b2:          -1 / 3,
		b3:          sqrt3 / 3,
		start_angle: 0,
	}
)

type Layout struct {
	size   points.Point
	origin points.Point

	// orientation can be flat or pointy hexes
	orientation Orientation

	// offset is the type of offset to use when converting between hex coordinates and screen coordinates.
	offset int
}

// NewLayout returns a flat layout with the given size and origin.
func NewLayout(size points.Point, origin points.Point, offset OffsetCoord_t) Layout {
	return Layout{
		orientation: layout_flat,
		size:        size,
		origin:      origin,
	}
}

func NewFlatEvenQLayout(size points.Point, origin points.Point) Layout {
	return Layout{
		orientation: layout_flat,
		size:        size,
		origin:      origin,
		offset:      EVEN,
	}
}

// --------------------------------------------------------------------------------------------------------------------
// hex to screen

func hex_to_pixel(layout Layout, h Hex) points.Point {
	var M = layout.orientation
	var x = (M.f0*float64(h.q) + M.f1*float64(h.r)) * layout.size.X
	var y = (M.f2*float64(h.q) + M.f3*float64(h.r)) * layout.size.Y
	return points.Point{X: x + layout.origin.X, Y: y + layout.origin.Y}
}

// --------------------------------------------------------------------------------------------------------------------
// screen to hex

func pixel_to_hex(layout Layout, p points.Point) FractionalHex {
	var M = layout.orientation
	var pt = points.Point{X: (p.X - layout.origin.X) / layout.size.X, Y: (p.Y - layout.origin.Y) / layout.size.Y}
	var q = M.b0*pt.X + M.b1*pt.Y
	var r = M.b2*pt.X + M.b3*pt.Y
	return FractionalHex{q: q, r: r, s: -q - r}
}

// --------------------------------------------------------------------------------------------------------------------
// drawing hex on screen

func hex_corner_offset(layout Layout, corner int) points.Point {
	size := layout.size
	angle := 2.0 * math.Pi * (layout.orientation.start_angle + float64(corner)) / 6
	return points.Point{X: size.X * math.Cos(angle), Y: size.Y * math.Sin(angle)}
}

func polygon_corners(layout Layout, h Hex) [6]points.Point {
	var corners [6]points.Point
	center := hex_to_pixel(layout, h)
	for i := 0; i < 6; i++ {
		offset := hex_corner_offset(layout, i)
		corners[i] = points.Point{X: center.X + offset.X, Y: center.Y + offset.Y}
	}
	return corners
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
	a_nudge := FractionalHex{float64(a.q) + 1e-6, float64(a.r) + 1e-6, float64(a.s) - 2e-6}
	b_nudge := FractionalHex{float64(b.q) + 1e-6, float64(b.r) + 1e-6, float64(b.s) - 2e-6}
	step := 1.0 / float64(max(N, 1))
	for i := 0; i <= N; i++ {
		results = append(results, hex_round(fractional_hex_lerp(a_nudge, b_nudge, float64(i)*step)))
	}
	return results
}

// --------------------------------------------------------------------------------------------------------------------
// rotation

func hex_rotate_left(a Hex) Hex {
	return Hex{q: -a.s, r: -a.q, s: -a.r}
}

func hex_rotate_right(a Hex) Hex {
	return Hex{q: -a.r, r: -a.s, s: -a.q}
}

// --------------------------------------------------------------------------------------------------------------------
// offset coordinates

type OffsetCoord struct {
	col, row int
}

type OffsetCoord_t bool

const (
	EvenOffset OffsetCoord_t = true
	OddOffset  OffsetCoord_t = false
)

const EVEN = +1

const ODD = -1

func qoffset_from_cube(offset int, h Hex) OffsetCoord {
	if !(offset == EVEN || offset == ODD) {
		panic("assert(offset == EVEN || offset == ODD)")
	}
	col := h.q
	row := h.r + int((h.q+offset*(h.q&1))/2)
	return OffsetCoord{col: col, row: row}
}

func qoffset_to_cube(offset int, h OffsetCoord) Hex {
	if !(offset == EVEN || offset == ODD) {
		panic("assert(offset == EVEN || offset == ODD)")
	}
	q := h.col
	r := h.row - int((h.col+offset*(h.col&1))/2)
	s := -q - r
	return Hex{q: q, r: r, s: s}
}

func roffset_from_cube(offset int, h Hex) OffsetCoord {
	if !(offset == EVEN || offset == ODD) {
		panic("assert(offset == EVEN || offset == ODD)")
	}
	col := h.q + int((h.r+offset*(h.r&1))/2)
	row := h.r
	return OffsetCoord{col: col, row: row}
}

func roffset_to_cube(offset int, h OffsetCoord) Hex {
	if !(offset == EVEN || offset == ODD) {
		panic("assert(offset == EVEN || offset == ODD)")
	}
	q := h.col - int((h.row+offset*(h.row&1))/2)
	r := h.row
	s := -q - r
	return Hex{q: q, r: r, s: s}
}
