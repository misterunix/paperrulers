package main

import (
	"fmt"
	"image/color"
)

func DrawLine(a Point, b Point, width float64, linecolor color.RGBA) {
	CreateGC()
	fmt.Println(a, b)
	Opt.GC.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	Opt.GC.SetLineWidth(width)
	Opt.GC.MoveTo(a.X, a.Y)
	Opt.GC.LineTo(b.X, b.Y)
	Opt.GC.Close()
	Opt.GC.FillStroke()
}
