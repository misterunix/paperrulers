package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"image/color"
)

func main() {

	style := 0

	flag.IntVar(&style, "style", 0, "page style\n  0 - lines\n  1 - dots\n  2 - cursive grid")

	flag.Float64Var(&Opt.Spacing, "sp", 7.0, "spacing between dots or lines in mm")
	flag.BoolVar(&Opt.Centermark, "c", false, "draw center dot or line")
	flag.StringVar(&Opt.PaperOrientation, "o", "L", "paper orientation. L for landscape, P for portrait")
	flag.StringVar(&Opt.PaperSize, "ps", "Letter", "paper size. Letter, A4, etc")
	flag.Float64Var(&Opt.Cursiveunits, "u", 5.0, "units for cursive grid, overrides spacing")
	flag.Float64Var(&Opt.Angle, "a", 0.0, "angle in degrees offset of center mark")
	flag.BoolVar(&Opt.Ladder, "l", false, "for blackletter 2/4/2/4")
	flag.BoolVar(&Opt.Dark, "dark", false, "dark lines")
	flag.StringVar(&Opt.PDFDir, "pdf", "~/tmp/paperrulers", "directory to save PDF files")
	flag.Parse()

	ensureDir("~/tmp/paperrulers")

	Opt.PaperSize = strings.ToLower(Opt.PaperSize)
	Opt.PaperOrientation = strings.ToLower(Opt.PaperOrientation)

	Opt.LRmargin = 12.7

	if Opt.Centermark {
		Opt.CenterSpaceing = Opt.Spacing / 2.0
	}

	Opt.LineWidth = 0.2 // line width in mm
	if Opt.PaperSize == "" {
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	switch Opt.PaperSize {
	case "letter":
		switch Opt.PaperOrientation {
		case "l":
			Opt.PageWidth = 279.4
			Opt.PageHeight = 215.9
			Opt.Margins = Opt.LRmargin / 2
		case "p":
			Opt.PageWidth = 215.9
			Opt.PageHeight = 279.4
			Opt.Margins = Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "a4":
		switch Opt.PaperOrientation {
		case "l":
			Opt.PageWidth = 297
			Opt.PageHeight = 210
			Opt.Margins = Opt.LRmargin
		case "p":
			Opt.PageWidth = 210
			Opt.PageHeight = 297
			Opt.Margins = Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}
	case "b5":
		switch Opt.PaperOrientation {
		case "l":
			Opt.PageWidth = 250
			Opt.PageHeight = 176
			Opt.Margins = Opt.LRmargin
		case "p":
			Opt.PageWidth = 176
			Opt.PageHeight = 250
			Opt.Margins = Opt.LRmargin
		default:
			fmt.Println("Invalid paper orientation")
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid paper size")
		os.Exit(1)
	}

	Opt.DarkBlack = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	if Opt.Dark {
		Opt.LightGray = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	} else {
		Opt.LightGray = color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff}
	}

	// os.Mkdir("pdf", 0755)
	// if Opt.dot {
	// 	if Opt.Centermark {
	// 		Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%d-center.pdf", Opt.PaperSize, Opt.PaperOrientation, int(Opt.Spacing))
	// 	} else {
	// 		Opt.Filename = fmt.Sprintf("pdf/dots-%s-%s-%d.pdf", Opt.PaperSize, Opt.PaperOrientation, int(Opt.Spacing))
	// 	}
	// } else {
	// 	if Opt.Centermark {
	// 		Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%d-center.pdf", Opt.PaperSize, Opt.PaperOrientation, int(Opt.Spacing))
	// 	} else {
	// 		Opt.Filename = fmt.Sprintf("pdf/lines-%s-%s-%d.pdf", Opt.PaperSize, Opt.PaperOrientation, int(Opt.Spacing))
	// 	}
	// }

	Opt.PageMarginLeft = Opt.Margins
	Opt.PageMarginRight = Opt.PageWidth - Opt.Margins
	Opt.PageMarginTop = Opt.Margins
	Opt.PageMarginBottom = Opt.PageHeight - Opt.Margins

	fmt.Printf("%v\n", Opt)

	switch style {
	case 0:
		DrawLines()
	case 1:
		DrawDots()
	case 2:
		CursiveGrid()
	default:
		fmt.Println("Invalid style")
		os.Exit(1)
	}

}

// check if directory exists and create it if not
func ensureDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dirName, err)
			os.Exit(1)
		}
	}
}
