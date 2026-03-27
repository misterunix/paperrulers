package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dpdf"
)

// createPDFBase creates a new PDF surface with a given orientation and a given unit
func CreatePDFBase() {
	//Opt.NewDest := draw2dpdf.NewPdf(Opt.PaperOrientation, "mm", Opt.PaperSize)
	Opt.Dest = draw2dpdf.NewPdf(Opt.PaperOrientation, "mm", Opt.PaperSize)
}

// create a new Graphic context
func CreateGC() {
	Opt.GC = draw2dpdf.NewGraphicContext(Opt.Dest)
}

func DrawLines() {
	if Opt.Centermark {
		Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f-center.pdf", Opt.PaperSize, Opt.PaperOrientation, Opt.Spacing)
	} else {
		Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%02.3f.pdf", Opt.PaperSize, Opt.PaperOrientation, Opt.Spacing)
	}

	CreatePDFBase()

	if Opt.Ladder {
		drawLadderLineGroup()
	}

	if Opt.Centermark {
		for y := Opt.PageMarginTop + Opt.CenterSpaceing; y <= Opt.PageMarginBottom; y += Opt.Spacing {
			DrawLine(Point{X: Opt.PageMarginLeft, Y: y},
				Point{X: Opt.PageMarginRight, Y: y},
				Opt.LineWidth, Opt.LightGray)
		}
	}
	draw2dpdf.SaveToPdfFile(Opt.Filename, Opt.Dest)

}

func drawDot(a Point, radius float64, width float64, linecolor color.RGBA) {
	CreateGC()
	Opt.GC.SetStrokeColor(color.RGBA{R: linecolor.R, G: linecolor.G, B: linecolor.B, A: linecolor.A})
	Opt.GC.SetLineWidth(width)
	Opt.GC.MoveTo(a.X, a.Y)
	Opt.GC.ArcTo(a.X, a.Y, radius, radius, 0, 2*math.Pi)
	Opt.GC.Close()
	Opt.GC.FillStroke()
}

func DrawDots() {
	if Opt.Centermark {
		Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%f-center.pdf", Opt.PaperSize, Opt.PaperOrientation, Opt.Spacing)
	} else {
		Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%f.pdf", Opt.PaperSize, Opt.PaperOrientation, Opt.Spacing)
	}

	CreatePDFBase()

	for y := Opt.PageMarginTop; y <= Opt.PageMarginBottom; y += Opt.Spacing {
		for x := Opt.PageMarginLeft; x <= Opt.PageMarginRight; x += Opt.Spacing {
			drawDot(Point{X: x, Y: y}, 0.15, Opt.LineWidth, Opt.DarkBlack)
		}
		//drawLine(Point{Opt.PageMarginLeft, y}, Point{Opt.PageMarginRight, y}, Opt.LineWidth, Opt.DarkBlack)
	}

	if Opt.Centermark {
		for y := Opt.PageMarginTop + Opt.CenterSpaceing; y <= Opt.PageMarginBottom; y += Opt.Spacing {
			for x := Opt.PageMarginLeft + Opt.CenterSpaceing; x <= Opt.PageMarginRight; x += Opt.Spacing {
				//angle := 0.0 //-12.0
				s := math.Sin(Opt.Angle*math.Pi/180) * Opt.Spacing
				drawDot(Point{X: x + s, Y: y}, 0.15, Opt.LineWidth, Opt.LightGray)
			}
		}
	}

	/*

		createGC()
		// set stroke color
		Opt.GC.SetStrokeColor(color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff})
		Opt.GC.SetLineWidth(Opt.LineWidth)

		Opt.border = 25 / 2.0
		xb := 25.0 - Opt.border
		yb := 25.0 - Opt.border

		for x := xb; x < Opt.PageWidth-Opt.border; x += Opt.Spacing {
			for y := yb; y < Opt.PageHeight-Opt.border; y += Opt.Spacing {

				Opt.GC.MoveTo(x, y)
				Opt.GC.ArcTo(x, y, 0.15, 0.15, 0, 2*math.Pi)
				Opt.GC.Close()
			}
		}

		Opt.GC.Close()
		Opt.GC.FillStroke()

		// center line if set
		if Opt.Centermark {

			Opt.GC.SetStrokeColor(color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xff})
			Opt.GC.SetLineWidth(Opt.LineWidth)
			for x := xb + Opt.Spacing; x < Opt.PageWidth-Opt.border; x += Opt.Spacing {
				for y := yb; y < Opt.PageHeight-Opt.border; y += Opt.Spacing {

					Opt.GC.MoveTo(x-(Opt.Spacing/2), y-(Opt.Spacing/2))
					Opt.GC.ArcTo(x-(Opt.Spacing/2), y-(Opt.Spacing/2), 0.15, 0.15, 0, 2*math.Pi)
					Opt.GC.Close()
				}
			}

			Opt.GC.Close()
			Opt.GC.FillStroke()
		}
	*/

	draw2dpdf.SaveToPdfFile(Opt.Filename, Opt.Dest)

}

func CursiveGrid() {

	Opt.Spacing = Opt.Cursiveunits
	Opt.Filename = fmt.Sprintf("pdf/cursive-%s-%s-%f-center.pdf", Opt.PaperSize, Opt.PaperOrientation, Opt.Spacing)

	CreatePDFBase()
	CreateGC()

	Opt.GC.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	Opt.GC.SetLineWidth(Opt.LineWidth)

	Opt.GC.SetStrokeColor(color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xff})

	// set line width
	Opt.GC.SetLineWidth(Opt.LineWidth)

	down := 0.0
	pos := 0.0

	for {

		// ascender line
		Opt.GC.SetLineWidth(0.5)
		Opt.GC.SetStrokeColor(Opt.DarkBlack)
		down = Opt.PageMarginTop + pos
		Opt.GC.MoveTo(Opt.PageMarginLeft, down)
		Opt.GC.LineTo(Opt.PageMarginRight, down)
		Opt.GC.Close()
		Opt.GC.FillStroke()

		// t-d line
		Opt.GC.SetLineWidth(0.2)
		Opt.GC.SetStrokeColor(Opt.LightGray)
		down = Opt.PageMarginTop + pos + Opt.Spacing

		Opt.GC.MoveTo(Opt.PageMarginLeft, down)
		Opt.GC.LineTo(Opt.PageMarginRight, down)

		Opt.GC.Close()
		Opt.GC.FillStroke()

		// x-height line
		Opt.GC.SetLineWidth(0.2)
		down = Opt.PageMarginTop + pos + (Opt.Spacing * 2)

		Opt.GC.MoveTo(Opt.PageMarginLeft, down)
		Opt.GC.LineTo(Opt.PageMarginRight, down)

		Opt.GC.Close()
		Opt.GC.FillStroke()

		// base line
		Opt.GC.SetLineWidth(0.5)
		Opt.GC.SetStrokeColor(Opt.DarkBlack)
		down = Opt.PageMarginTop + pos + (Opt.Spacing * 3)

		Opt.GC.MoveTo(Opt.PageMarginLeft, down)
		Opt.GC.LineTo(Opt.PageMarginRight, down)

		Opt.GC.Close()
		Opt.GC.FillStroke()

		// descender line
		Opt.GC.SetLineWidth(0.5)
		Opt.GC.SetStrokeColor(Opt.DarkBlack)
		down = Opt.PageMarginTop + pos + (Opt.Spacing * 5)

		Opt.GC.MoveTo(Opt.PageMarginLeft, down)
		Opt.GC.LineTo(Opt.PageMarginRight, down)

		Opt.GC.Close()
		Opt.GC.FillStroke()

		pos += (Opt.Spacing * 6)

		//fmt.Println(pos, (Opt.PageHeight - Opt.PageMarginBottom))

		if pos > (Opt.PageHeight - Opt.Margins) {
			break
		}

	}

	draw2dpdf.SaveToPdfFile(Opt.Filename, Opt.Dest)
}
