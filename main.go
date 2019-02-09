package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, heigh=%v \n", w, h)
	pdf.AddPage()

	// Basic Text Stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 30)
	_, lineheight := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineheight, "Hello, Test")
	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineheight = pdf.GetFontSize()
	pdf.MoveTo(0, lineheight*2.0)
	pdf.MultiCell(0, lineheight*1.5, "Here is some text. If it is too long it will be word wrapped automatically. If there is a new line it will be \n wrapped as well (unlike other ways of writing text in gofpdf).", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Basic shapes
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{110, 250},
		{160, 300},
		{110, 350},
		{60, 300},
	}, "F")

	pdf.Polygon([]gofpdf.PointType{
		{110, 350},
		{160, 400},
		{110, 450},
		{60, 400},
	}, "F")

	pdf.ImageOptions("images/jump.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")
	drawGrid(pdf)
	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		log.Fatalln(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineht := pdf.GetFontSize()
		pdf.Text(x, lineht, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
