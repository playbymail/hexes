// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/playbymail/hexes"
	"log"
	"math"
)

func main() {
	cx, cy := 512.0, 512.0

	layout := hexes.NewFlatEvenLayout(hexes.NewPoint(50, 50), hexes.NewPoint(cx, cy))
	err := drawHexes("flat-even.png", newHexMap(5), layout)
	if err != nil {
		log.Fatal(err)
	}

	layout = hexes.NewPointyEvenLayout(hexes.NewPoint(50, 50), hexes.NewPoint(cx, cy))
	err = drawHexes("pointy-even.png", newHexMap(5), layout)
	if err != nil {
		log.Fatal(err)
	}

	layout = hexes.NewFlatEvenLayout(hexes.NewPoint(50, 50), hexes.NewPoint(100, 100))
	err = drawHexes("flat-even-square.png", newSquareMap(layout, 5), layout)
	if err != nil {
		log.Fatal(err)
	}

	layout = hexes.NewPointyEvenLayout(hexes.NewPoint(50, 50), hexes.NewPoint(100, 100))
	err = drawHexes("pointy-even-square.png", newSquareMap(layout, 5), layout)
	if err != nil {
		log.Fatal(err)
	}

	layout = hexes.NewFlatOddLayout(hexes.NewPoint(50, 50), hexes.NewPoint(100, 100))
	err = drawHexes("flat-odd-square.png", newSquareMap(layout, 5), layout)
	if err != nil {
		log.Fatal(err)
	}

	err = drawHexesOnImage("testmap.png", "hex-map.png")
	if err != nil {
		log.Fatal(err)
	}

	err = drawHexesOnScaledImage("testmap.png", "hex-scaled-map.png")
	if err != nil {
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

func newSquareMap(l hexes.Layout, n int) (h []hexes.Hex) {
	for column := 0; column <= n; column++ {
		for row := 0; row <= n; row++ {
			h = append(h, l.OffsetToHex(column, row))
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

		column, row := l.HexToOffset(h)
		dc.DrawStringAnchored(fmt.Sprintf("%d, %d", column, row), cp.X, cp.Y, 0.5, 0.5)

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

func drawHexesOnImage(input, output string) error {
	// Load the PNG image from the file
	img, err := gg.LoadImage(input)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to load image", err))
	}

	// Create a new drawing context with the size of the image
	dc := gg.NewContextForImage(img)

	// Get the width and height of the image
	mapWidth, mapHeight := float64(dc.Width()), float64(dc.Height())
	log.Printf("hexes: map width  %8.2f x %8.2f\n", mapWidth, mapHeight)
	cx, cy := mapWidth/2, mapHeight/2
	log.Printf("hexes: map center %8.2f x %8.2f\n", cx, cy)

	// we want to get about 100 hexes onto the map
	hexWidth, hexHeight := 12.5, 12.5 //(mapWidth - 10) / 100
	hexesWide := int(math.Ceil(mapWidth / hexWidth))
	log.Printf("hexes: hex width  %8.2f x %8.2f\n", hexWidth, hexHeight)
	log.Printf("hexes: hexes wide %8d\n", hexesWide)

	layout := hexes.NewFlatOddLayout(hexes.NewPoint(hexWidth, hexHeight), hexes.NewPoint(float64(hexWidth)/2, float64(hexHeight)/2))

	hv := newSquareMap(layout, int(math.Ceil(mapWidth/hexWidth)))

	// draw them
	for _, h := range hv {
		_, corners := layout.Points(h)
		pt := corners[5]
		x1, y1 := pt.X, pt.Y
		for i := 0; i < 6; i++ {
			pt = corners[i]
			x2, y2 := pt.X, pt.Y
			dc.SetRGB(25, 25, 25)
			dc.SetLineWidth(1)
			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
			x1, y1 = x2, y2
		}
	}

	// Save the modified image as a new PNG file
	if err := dc.SavePNG(output); err != nil {
		return errors.Join(fmt.Errorf("failed to save image", err))
	}

	return nil
}

func drawHexesOnScaledImage(input, output string) error {
	// Load the PNG image from the file
	img, err := gg.LoadImage(input)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to load image", err))
	}

	// Get the original width and height of the image
	origWidth, origHeight := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
	log.Printf("hexes: map width  %8.2f x %8.2f\n", origWidth, origHeight)

	// Calculate new dimensions (25% wider and taller)
	newWidth, newHeight := origWidth*1.25, origHeight*1.25
	log.Printf("hexes: map width  %8.2f x %8.2f\n", newWidth, newHeight)

	// Create a new drawing context with the new size
	dc := gg.NewContext(int(newWidth), int(newHeight))

	// Scale the drawing context
	dc.Scale(1.25, 1.25)

	// Draw the original image on the scaled context
	dc.DrawImage(img, 0, 0)

	// Reset the scaling for further drawing operations
	dc.Identity()

	// Get the width and height of the image
	mapWidth, mapHeight := float64(dc.Width()), float64(dc.Height())
	log.Printf("hexes: map width  %8.2f x %8.2f\n", mapWidth, mapHeight)
	cx, cy := mapWidth/2, mapHeight/2
	log.Printf("hexes: map center %8.2f x %8.2f\n", cx, cy)

	// we want to get about 100 hexes onto the map
	hexWidth, hexHeight := 12.5, 12.5 //(mapWidth - 10) / 100
	hexesWide := int(math.Ceil(mapWidth / hexWidth))
	log.Printf("hexes: hex width  %8.2f x %8.2f\n", hexWidth, hexHeight)
	log.Printf("hexes: hexes wide %8d\n", hexesWide)

	layout := hexes.NewFlatOddLayout(hexes.NewPoint(hexWidth, hexHeight), hexes.NewPoint(float64(hexWidth)/2, float64(hexHeight)/2))

	hv := newSquareMap(layout, int(math.Ceil(mapWidth/hexWidth)))

	// draw them
	for _, h := range hv {
		_, corners := layout.Points(h)
		pt := corners[5]
		x1, y1 := pt.X, pt.Y
		for i := 0; i < 6; i++ {
			pt = corners[i]
			x2, y2 := pt.X, pt.Y
			dc.SetRGB(25, 25, 25)
			dc.SetLineWidth(1)
			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
			x1, y1 = x2, y2
		}
	}

	// Save the modified image as a new PNG file
	if err := dc.SavePNG(output); err != nil {
		return errors.Join(fmt.Errorf("failed to save image", err))
	}

	return nil
}
