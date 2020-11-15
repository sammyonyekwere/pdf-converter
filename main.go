package main

import (
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	// creating new report PDF
	pdf := newReport()

	// pdf table header
	pdf = header(pdf, []string{"1st column", "2nd", "3rd", "4th", "5th", "6th"})

	// pdf table content
	pdf = table(pdf)

	if pdf.Err() {
		log.Fatalf("failed! %s", pdf.Error())
	}

	err := pdf.OutputFileAndClose("report-example.pdf")
	if err != nil {
		log.Fatalf("error saving pdf file: %s", err)
	}
}

func newReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "PDF Report")
	pdf.Ln(12)

	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)

	return pdf

}

func header(pdf *gofpdf.Fpdf, headerText []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 20)
	pdf.SetFillColor(240, 240, 240)

	for _, str := range headerText {
		pdf.CellFormat(40, 10, str, "1", 0, "", true, 0, "")
	}

	pdf.Ln(-1)

	return pdf
}

func table(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {

	pdf.SetFont("Times", "", 10)
	pdf.SetFillColor(255, 255, 255)

	pdf.CellFormat(40, 10, "text", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 10, "same", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")

	pdf.Ln(-1)

	return pdf
}
