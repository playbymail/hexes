// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/fogleman/gg"
	"github.com/playbymail/hexes"
	"log"
)

func main() {
	cx, cy := 512.0, 512.0

	hv := newHexMap(5)

	if err := drawHexes("flat-even.png", hv, hexes.NewFlatEvenLayout(hexes.NewPoint(50, 50), hexes.NewPoint(cx, cy))); err != nil {
		log.Fatal(err)
	}

	if err := drawHexes("pointy-odd.png", hv, hexes.NewPointyOddLayout(hexes.NewPoint(50, 50), hexes.NewPoint(cx, cy))); err != nil {
		log.Fatal(err)
	}
}

func newHexMap(n int) (h []hexes.Hex) {
	for q := -n; q <= n; q++ {
		r1, r2 := max(-n, -q-n), min(n, -q+n)
		for r := r1; r <= r2; r++ {
			h = append(h, hexes.NewHex(q, r, -q-r))
		}
	}
	return h
}

func drawHexes(path string, hv []hexes.Hex, l hexes.Layout) error {
	// create a new gg image
	dc := gg.NewContext(1024, 1024)

	// draw them
	for _, h := range hv {
		cp, corners := l.Points(h)
		dc.DrawCircle(cp.X, cp.Y, 3)
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		pt := corners[5]
		x1, y1 := pt.X, pt.Y
		for i := 0; i < 6; i++ {
			pt = corners[i]
			x2, y2 := pt.X, pt.Y
			dc.SetRGB(25, 25, 25)
			dc.SetLineWidth(2)
			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
			x1, y1 = x2, y2
		}
	}

	err := dc.SavePNG(path)
	if err == nil {
		log.Printf("created %s\n", path)
	}

	return err
}
