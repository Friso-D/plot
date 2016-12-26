// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"log"
	"math"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/internal/cmpimg"
	"github.com/gonum/plot/palette/moreland"
)

// ExamplePolygon_holes draws a polygon with holes. The output of this
// example is at
// https://github.com/gonum/plot/blob/master/plotter/testdata/polygon_holes_golden.png.
func ExamplePolygon_holes() {
	// Create an outer ring.
	outer1 := XYs{{X: 0, Y: 0}, {X: 4, Y: 0}, {X: 4, Y: 4}, {X: 0, Y: 4}}

	// create an inner ring with the same
	// winding order as the outer ring.
	inner1 := XYs{{X: 0.5, Y: 0.5}, {X: 1.5, Y: 0.5}, {X: 1.5, Y: 1.5}, {X: 0.5, Y: 1.5}}

	// create an inner polygon with the opposite
	// winding order as the outer polygon.
	inner2 := XYs{{X: 3.5, Y: 2.5}, {X: 2.5, Y: 2.5}, {X: 2.5, Y: 3.5}, {X: 3.5, Y: 3.5}}

	poly, err := NewPolygon(outer1, inner1, inner2)
	if err != nil {
		log.Panic(err)
	}
	poly.Color = color.NRGBA{B: 255, A: 255}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Polygon with holes"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(poly)

	err = p.Save(100, 100, "testdata/polygon_holes.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestPolygon_holes(t *testing.T) {
	cmpimg.CheckPlot(ExamplePolygon_holes, t, "polygon_holes.png")
}

// ExamplePolygon_hexagons creates a heat map with hexagon shapes.
// The output of this example is at
// https://github.com/gonum/plot/blob/master/plotter/testdata/polygon_hexagons_golden.png.
func ExamplePolygon_hexagons() {
	// hex returns a hexagon centered at (x,y) with radius r.
	hex := func(x, y, r float64) XYs {
		g := make(XYs, 6)
		for i := 0; i < 6; i++ {
			g[i].X = x + r*math.Cos(math.Pi*2/6*float64(i))
			g[i].Y = y + r*math.Sin(math.Pi*2/6*float64(i))
		}
		return g
	}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Hexagons"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	colorMap := moreland.SmoothBlueRed()
	colorMap.SetMax(2)
	colorMap.SetMin(-2)
	colorMap.SetConvergePoint(0)

	const (
		r = math.Pi / 4 // r is the hexagon radius
		// x0 and y0 are the beginning coordinates for the hexagon plot
		x0 = 0.0
		y0 = 0.0
		// nx and ny are the number of hexagons in the x and y directions.
		nx = 5
		ny = 5
	)
	// dx and dy are the distance between hexgons
	dx := 3 * r
	dy := r * math.Sqrt(3)

	xstart := []float64{x0, x0 - 1.5*r}
	ystart := []float64{y0, y0 - r}
	for i, xmin := range xstart {
		ymin := ystart[i]
		x := xmin
		for ix := 0; ix < nx; ix++ {
			y := ymin
			for iy := 0; iy < ny; iy++ {
				var poly *Polygon
				poly, err = NewPolygon(hex(x, y, r))
				if err != nil {
					log.Panic(err)
				}
				poly.Color, err = colorMap.At(math.Sin(x) + math.Sin(y))
				if err != nil {
					log.Panic(err)
				}
				poly.LineStyle.Width = 0
				p.Add(poly)
				y += dy
			}
			x += dx
		}
	}
	if err = p.Save(100, 100, "testdata/polygon_hexagons.png"); err != nil {
		log.Panic(err)
	}
}

func TestPolygon_hexagons(t *testing.T) {
	cmpimg.CheckPlot(ExamplePolygon_hexagons, t, "polygon_hexagons.png")
}
