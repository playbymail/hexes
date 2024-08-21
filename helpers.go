// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package hexes

//// --------------------------------------------------------------------------------------------------------------------
//// geometry is from https://www.redblobgames.com/grids/hexagons/#basics
////
//// A regular hexagon can be inscribed within two circles. The first is the outer circle (circumcircle?),
//// which is the circle that passes through all six corners of the hexagon. The second is the inner circle,
//// which is the circle that passes through the centers of all six sides of the hexagon.
////
//// A pointy hexagon is a flat hexagon rotated 30 degrees.
////
//// The source material assumes the Y-axis is pointing down, the X-axis is pointing to the right,
//// and angles increase clockwise. This seems to be setup to reflect screen-coordinates.
//// We may have to adjust if our Y-axis is pointing up and angles increase counter-clockwise.
//
//// --------------------------------------------------------------------------------------------------------------------
//// spacing is from https://www.redblobgames.com/grids/hexagons/#spacing
//
//// If it is a flat hexagon, the size is the radius of the inner circle. The width and height are
//// sqrt(3) * size and 2 * size, respectively.
//
//func flat_hex_spacing(size float64) (horizontal, vertical float64) {
//	horizontal = size * 3 / 2
//	vertical = size * sqrt3
//	return horizontal, vertical
//}
//
//// If we have a "pointy" hexagon, its "size" is the radius of the outer circle. The width and height
//// are 2 * size and sqrt(3) * size, respectively.
//
//func pointy_hex_spacing(size float64) (horizontal, vertical float64) {
//	horizontal = size * sqrt3
//	vertical = size * 3 / 2
//	return horizontal, vertical
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// angles are from https://www.redblobgames.com/grids/hexagons/#angles
//
//// flat_hex_corner returns the coordinates of the corner of a hexagon centered at the given point.
//func flat_hex_corner(center Point, size float64, i int) Point {
//	var angle_deg = 60 * float64(i)
//	var angle_rad = math.Pi / 180 * angle_deg
//	return Point{X: center.X + size*math.Cos(angle_rad), Y: center.Y + size*math.Sin(angle_rad)}
//}
//
//// flat_hex_vertices returns the coordinates of the corners of a hexagon centered at the given point.
//func flat_hex_vertices(center Point, size float64) [6]Point {
//	var vertices [6]Point
//	for i := 0; i < 6; i++ {
//		vertices[i] = flat_hex_corner(center, size, i)
//	}
//	return vertices
//}
//
//// pointy_hex_corner returns the coordinates of the corner of a hexagon centered at the given point.
//// The direction is an integer from 0 to 5, with 0 being "due east" and 5 being "northeast."
//func pointy_hex_corner(center Point, size float64, i int) Point {
//	var angle_deg = float64(60*i) - 30.0
//	var angle_rad = math.Pi / 180.0 * angle_deg
//	return Point{X: center.X + size*math.Cos(angle_rad), Y: center.Y + size*math.Sin(angle_rad)}
//}
//
//// pointy_hex_vertices returns the coordinates of the corners of a hexagon centered at the given point.
//func pointy_hex_vertices(center Point, size float64) [6]Point {
//	var vertices [6]Point
//	for i := 0; i < 6; i++ {
//		vertices[i] = pointy_hex_corner(center, size, i)
//	}
//	return vertices
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// offset coordinates are from https://www.redblobgames.com/grids/hexagons/#coordinates-offset
//
//type OffsetCoord_t struct {
//	col, row int
//}
//
//// OddR_t implements an immutable, regular hexagon using offset coordinates for a
//// horizontal layout (pointy tops) and odd rows shoved right.
//type OddR_t struct {
//	col, row int
//}
//
//// EvenR_t implements an immutable, regular hexagon using offset coordinates for a
//// horizontal layout (pointy tops) and even rows shoved right.
//type EvenR_t struct {
//	col, row int
//}
//
//// OddQ_t implements an immutable, regular hexagon using offset coordinates for a
//// vertical layout (flat tops) and odd columns shoved down.
//type OddQ_t struct {
//	col, row int
//}
//
//// EvenQ_t implements an immutable, regular hexagon using offset coordinates for a
//// vertical layout (flat tops) and even columns shoved down.
//type EvenQ_t struct {
//	col, row int
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// cube coordinates are from https://www.redblobgames.com/grids/hexagons/#coordinates-cube
//
//type Cube_t struct {
//	q, r, s int
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// conversions are from https://www.redblobgames.com/grids/hexagons/#conversions-offset
//
//// convert cube coordinates to offset coordinates (pointy-top, push odd rows down)
//func cube_to_oddr(hex Cube_t) OddR_t {
//	var col = hex.q + (hex.r-(hex.r&1))/2
//	var row = hex.r
//	return OddR_t{col: col, row: row}
//}
//
//// convert offset coordinates to cube coordinates  (pointy-top, push odd rows down)
//func oddr_to_cube(hex OddR_t) Cube_t {
//	var q = hex.col - (hex.row-(hex.row&1))/2
//	var r = hex.row
//	return Cube_t{q: q, r: r, s: -q - r}
//}
//
//// convert cube coordinates to offset coordinates (pointy-top, push even rows down)
//func cube_to_evenr(hex Cube_t) EvenR_t {
//	var col = hex.q + (hex.r+(hex.r&1))/2
//	var row = hex.r
//	return EvenR_t{col: col, row: row}
//}
//
//// convert offset coordinates to cube coordinates  (pointy-top, push even rows down)
//func evenr_to_cube(hex EvenR_t) Cube_t {
//	var q = hex.col - (hex.row+(hex.row&1))/2
//	var r = hex.row
//	return Cube_t{q: q, r: r, s: -q - r}
//}
//
//// convert cube coordinates to offset coordinates (flat-top, push odd columns right)
//func cube_to_oddq(hex Cube_t) OddQ_t {
//	var col = hex.q
//	var row = hex.r + (hex.q-(hex.q&1))/2
//	return OddQ_t{col: col, row: row}
//}
//
//// convert offset coordinates to cube coordinates (flat-top, push odd columns right)
//func oddq_to_cube(hex OddQ_t) Cube_t {
//	var q = hex.col
//	var r = hex.row - (hex.col-(hex.col&1))/2
//	return Cube_t{q: q, r: r, s: -q - r}
//}
//
//// convert cube coordinates to offset coordinates (flat-top, push even columns right)
//func cube_to_evenq(hex Cube_t) EvenQ_t {
//	var col = hex.q
//	var row = hex.r + (hex.q+(hex.q&1))/2
//	return EvenQ_t{col: col, row: row}
//}
//
//// convert offset coordinates to cube coordinates (flat-top, push even columns right)
//func evenq_to_cube(hex EvenQ_t) Cube_t {
//	var q = hex.col
//	var r = hex.row - (hex.col+(hex.col&1))/2
//	return Cube_t{q: q, r: r, s: -q - r}
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// neighbors are from https://www.redblobgames.com/grids/hexagons/#neighbors-cube
//
//// for flat hexes, the direction is counter-clockwise starting from the "southeast."
////   0 -> southeast              30 degrees
////   1 -> northeast             330 degrees
////   2 -> north                 270 degrees
////   3 -> northwest             210 degrees
////   4 -> southwest             150 degrees
////   5 -> south                  90 degrees
////
//// for pointy hexes, the direction is counter-clockwise starting from the "east."
////   0 -> east                    0 degrees
////   1 -> north by northeast    300 degrees
////   2 -> north by northwest    240 degrees
////   3 -> west                  180 degrees
////   4 -> south by southwest    120 degrees
////   5 -> south by southeast     60 degrees
//
//var cube_direction_vectors = [6]Cube_t{
//	{q: +1, r: 0, s: -1},
//	{q: +1, r: -1, s: 0},
//	{q: 0, r: -1, s: +1},
//	{q: -1, r: 0, s: +1},
//	{q: -1, r: +1, s: 0},
//	{q: 0, r: +1, s: -1},
//}
//
//func cube_direction(direction int) Cube_t {
//	return cube_direction_vectors[direction]
//}
//
//func cube_add(hex, vec Cube_t) Cube_t {
//	return Cube_t{hex.q + vec.q, hex.r + vec.r, hex.s + vec.s}
//}
//
//func cube_neighbor(cube Cube_t, direction int) Cube_t {
//	return cube_add(cube, cube_direction(direction))
//}
//
//var oddr_direction_differences = [2][6][2]int{
//	// even rows
//	{{+1, 0}, {0, -1}, {-1, -1}, {-1, 0}, {-1, +1}, {0, +1}},
//	// odd rows
//	{{+1, 0}, {+1, -1}, {0, -1}, {-1, 0}, {0, +1}, {+1, +1}},
//}
//
//func oddr_offset_neighbor(hex OddR_t, direction int) OddR_t {
//	var parity = hex.row & 1
//	var diff = oddr_direction_differences[parity][direction]
//	return OddR_t{col: hex.col + diff[0], row: hex.row + diff[1]}
//}
//
//var evenr_direction_differences = [2][6][2]int{
//	// even rows
//	{{+1, 0}, {+1, -1}, {0, -1}, {-1, 0}, {0, +1}, {+1, +1}},
//	// odd rows
//	{{+1, 0}, {0, -1}, {-1, -1}, {-1, 0}, {-1, +1}, {0, +1}},
//}
//
//func evenr_offset_neighbor(hex EvenR_t, direction int) EvenR_t {
//	var parity = hex.row & 1
//	var diff = evenr_direction_differences[parity][direction]
//	return EvenR_t{col: hex.col + diff[0], row: hex.row + diff[1]}
//}
//
//var oddq_direction_differences = [2][6][2]int{
//	// even cols
//	{{+1, 0}, {+1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {0, +1}},
//	// odd cols
//	{{+1, +1}, {+1, 0}, {0, -1}, {-1, 0}, {-1, +1}, {0, +1}},
//}
//
//func oddq_offset_neighbor(hex OddQ_t, direction int) OddQ_t {
//	var parity = hex.col & 1
//	var diff = oddq_direction_differences[parity][direction]
//	return OddQ_t{col: hex.col + diff[0], row: hex.row + diff[1]}
//}
//
//var evenq_direction_differences = [2][6][2]int{
//	// even cols
//	{{+1, +1}, {+1, 0}, {0, -1}, {-1, 0}, {-1, +1}, {0, +1}},
//	// odd cols
//	{{+1, 0}, {+1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {0, +1}},
//}
//
//func evenq_offset_neighbor(hex EvenQ_t, direction int) EvenQ_t {
//	var parity = hex.col & 1
//	var diff = evenq_direction_differences[parity][direction]
//	return EvenQ_t{col: hex.col + diff[0], row: hex.row + diff[1]}
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// diagonals are from https://www.redblobgames.com/grids/hexagons/#neighbors-diagonal
//
//var cube_diagonal_vectors = [6]Cube_t{
//	Cube_t{q: +2, r: -1, s: -1}, Cube_t{q: +1, r: -2, s: +1}, Cube_t{q: -1, r: -1, s: +2},
//	Cube_t{q: -2, r: +1, s: +1}, Cube_t{q: -1, r: +2, s: -1}, Cube_t{q: +1, r: +1, s: -2},
//}
//
//// cube_diagonal_neighbor returns the neighbor in the given direction, skipping out to two hexes.
////
//// for flat hexes, the direction is counter-clockwise starting from the "east."
////   0 -> east                        0 degrees
////   1 -> north by northeast        300 degrees
////   2 -> north by northwest        240 degrees
////   3 -> west                      180 degrees
////   4 -> south by southwest        120 degrees
////   5 -> south by southeast         60 degrees
////
//// for pointy hexes, the direction is counter-clockwise starting from the "northeast."
////   0 -> northeast                 330 degrees
////   1 -> north                     270 degrees
////   2 -> northwest                 210 degrees
////   3 -> southwest                 150 degrees
////   4 -> south                      90 degrees
////   5 -> southeast                  30 degrees
//
//func cube_diagonal_neighbor(cube Cube_t, direction int) Cube_t {
//	return cube_add(cube, cube_diagonal_vectors[direction])
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// distances are from https://www.redblobgames.com/grids/hexagons/#distances
////
//// distance is the number of hexes between two hexes. adjacent hexes are 1 unit away.
//
//// abs returns the absolute value of i.
//func abs(i int) int {
//	if i < 0 {
//		return -i
//	}
//	return i
//}
//
//func cube_subtract(a, b Cube_t) Cube_t {
//	return Cube_t{q: a.q - b.q, r: a.r - b.r, s: a.s - b.s}
//}
//
//func cube_distance(a, b Cube_t) int {
//	return max(abs(a.q-b.q), abs(a.r-b.r), abs(a.s-b.s))
//}
//
//func oddr_distance(a, b OddR_t) int {
//	return cube_distance(oddr_to_cube(a), oddr_to_cube(b))
//}
//
//func evenr_distance(a, b EvenR_t) int {
//	return cube_distance(evenr_to_cube(a), evenr_to_cube(b))
//}
//
//func oddq_distance(a, b OddQ_t) int {
//	return cube_distance(oddq_to_cube(a), oddq_to_cube(b))
//}
//
//func evenq_distance(a, b EvenQ_t) int {
//	return cube_distance(evenq_to_cube(a), evenq_to_cube(b))
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// line drawing is from https://www.redblobgames.com/grids/hexagons/#line-drawing
//
//// lerp is a linear interpolation function that is a helper for cube_lerp.
//func lerp(a, b, t float64) float64 {
//	return a + (b-a)*t
//}
//
//// cube_lerp is a linear interpolation function that is a helper for cube_linedraw.
//func cube_lerp(a, b Cube_t, t float64) (q, r, s float64) {
//	q = lerp(float64(a.q), float64(b.q), t)
//	r = lerp(float64(a.r), float64(b.r), t)
//	s = lerp(float64(a.s), float64(b.s), t)
//	return q, r, s
//}
//
//// cube_linedraw returns a list of hexes that lie on the line between a and b.
//func cube_linedraw(a, b Cube_t) (results []Cube_t) {
//	var N = cube_distance(a, b)
//	for i := 0; i <= N; i++ {
//		q, r, s := cube_lerp(a, b, 1.0/float64(N*i))
//		results = append(results, cube_round(q, r, s))
//	}
//	return results
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// reflections are from https://www.redblobgames.com/grids/hexagons/#reflection
//
//func reflectQ(h Cube_t) Cube_t {
//	return Cube_t{q: h.q, r: h.s, s: h.r}
//}
//
//func reflectR(h Cube_t) Cube_t {
//	return Cube_t{q: h.s, r: h.r, s: h.q}
//}
//
//func reflectS(h Cube_t) Cube_t {
//	return Cube_t{q: h.r, r: h.q, s: h.s}
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// hex to pixel is from https://www.redblobgames.com/grids/hexagons/#hex-to-pixel
//
//// pointy_hex_to_pixel returns the pixel coordinates of the center of the given hex.
//func pointy_hex_to_pixel(hex Cube_t, size float64) Point {
//	r, q := float64(hex.r), float64(hex.q)
//	return Point{X: size * (sqrt3*q + sqrt3/2*r), Y: size * (3.0 / 2.0 * r)}
//}
//
//// flat_hex_to_pixel returns the pixel coordinates of the center of the given hex.
//func flat_hex_to_pixel(hex Cube_t, size float64) Point {
//	r, q := float64(hex.r), float64(hex.q)
//	return Point{X: size * (3.0 / 2.0 * q), Y: size * (sqrt3/2*q + sqrt3*r)}
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// pixel to hex is from https://www.redblobgames.com/grids/hexagons/#pixel-to-hex
//
//func pixel_to_flat_hex(point Point, size float64) Cube_t {
//	var q = (2. / 3 * point.X) / size
//	var r = (-1./3*point.Y + sqrt3/3*point.Y) / size
//	var s = -q - r
//	return cube_round(q, r, s)
//}
//
//func pixel_to_pointy_hex(point Point, size float64) Cube_t {
//	var q = (sqrt3/3*point.X - 1./3*point.Y) / size
//	var r = (2. / 3 * point.Y) / size
//	var s = -q - r
//	return cube_round(q, r, s)
//}
//
//// --------------------------------------------------------------------------------------------------------------------
//// rounding is from https://www.redblobgames.com/grids/hexagons/#rounding
//
//func cube_round(fq, fr, fs float64) Cube_t {
//	var q = math.Round(fq)
//	var r = math.Round(fr)
//	var s = math.Round(fs)
//
//	var dq = math.Abs(q - fq)
//	var dr = math.Abs(r - fr)
//	var ds = math.Abs(s - fs)
//
//	if dq > dr && dq > ds {
//		q = -r - s
//	} else if dr > ds {
//		r = -q - s
//	} else {
//		s = -q - r
//	}
//
//	return Cube_t{q: int(q), r: int(r), s: int(s)}
//}
