package main

import (
	"image/color"
)

func drawLadder(x, y float64) {
	//fmt.Println(x, y)

	for i := 1; i <= 2; i = i + 2 { // y drop
		ym := float64(i) //+ Opt.Spacing

		CreateGC()
		Opt.GC.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.GC.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.GC.SetLineWidth(0.1)
		Opt.GC.BeginPath()
		bx := x
		by := y * ym
		Opt.GC.MoveTo(bx, by)                              // ul
		Opt.GC.LineTo(bx+Opt.Spacing, by)                  // ur
		Opt.GC.LineTo(bx+Opt.Spacing, (by+Opt.Spacing)*ym) // lr
		Opt.GC.LineTo(bx, (by+Opt.Spacing)*ym)             // lr
		Opt.GC.Close()
		Opt.GC.FillStroke()

		CreateGC()
		Opt.GC.SetFillColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.GC.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
		Opt.GC.SetLineWidth(0.1)
		Opt.GC.BeginPath()
		bx = x + Opt.Spacing
		by = (y + Opt.Spacing) * ym
		Opt.GC.MoveTo(bx, by)                         // ul
		Opt.GC.LineTo(bx+Opt.Spacing, by)             // ur
		Opt.GC.LineTo(bx+Opt.Spacing, by+Opt.Spacing) // lr
		Opt.GC.LineTo(bx, by+Opt.Spacing)             // lr
		Opt.GC.Close()
		Opt.GC.FillStroke()

	}

}

func drawLadderLines(yPos float64) {

	y := yPos
	// everything is in units of spacing

	// top ascender
	DrawLine(Point{X: Opt.PageMarginLeft, Y: y},
		Point{X: Opt.PageMarginRight, Y: y}, Opt.LineWidth, Opt.DarkBlack)

	// mid ascender
	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing},
		Opt.LineWidth,
		Opt.LightGray)

	drawLadder(Opt.PageMarginLeft, y)
	y = y + (Opt.Spacing * 2)

	// top x-height
	DrawLine(Point{X: Opt.PageMarginLeft, Y: y},
		Point{X: Opt.PageMarginRight, Y: y}, Opt.LineWidth, Opt.DarkBlack)

	// 1 down from x-height
	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing}, Opt.LineWidth,
		Opt.LightGray)

	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing*2},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing*2},
		Opt.LineWidth, Opt.LightGray)

	drawLadder(Opt.PageMarginLeft, y)

	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing}, Opt.LineWidth,
		Opt.LightGray)

	y = y + (Opt.Spacing * 2)
	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing}, Opt.LineWidth,
		Opt.LightGray)

	drawLadder(Opt.PageMarginLeft, y)

	y = y + (Opt.Spacing * 2)

	DrawLine(Point{X: Opt.PageMarginLeft, Y: y},
		Point{X: Opt.PageMarginRight, Y: y}, Opt.LineWidth,
		Opt.DarkBlack)

	DrawLine(Point{X: Opt.PageMarginLeft, Y: y + Opt.Spacing},
		Point{X: Opt.PageMarginRight, Y: y + Opt.Spacing}, Opt.LineWidth,
		Opt.LightGray)

	drawLadder(Opt.PageMarginLeft, y)

	y = y + (Opt.Spacing * 2)

	DrawLine(Point{X: Opt.PageMarginLeft, Y: y},
		Point{X: Opt.PageMarginRight, Y: y}, Opt.LineWidth, Opt.DarkBlack)
}

// drawLadderLines draws a group of ladder lines starting at the y position
// provided
func drawLadderLineGroup() {

	for y := Opt.PageMarginTop; y <= Opt.PageMarginBottom; y += (Opt.Spacing * 12) {
		if y+(Opt.Spacing*8) > Opt.PageMarginBottom {
			return
		}
		drawLadderLines(y)
	}

}
