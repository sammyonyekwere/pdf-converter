package main

import (
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "PDF Report")
	pdf.Ln(12)

	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)

	pdf.SetFont("Times", "", 20)
	pdf.SetFillColor(240, 240, 240)

	pdf.CellFormat(40, 10, "12", "1", 0, "", true, 0, "")
	pdf.CellFormat(40, 10, "22", "1", 0, "", true, 0, "")
	pdf.CellFormat(40, 10, "25", "1", 0, "", true, 0, "")
	pdf.CellFormat(40, 10, "33", "1", 0, "", true, 0, "")
	pdf.CellFormat(40, 10, "55", "1", 0, "", true, 0, "")
	pdf.CellFormat(40, 10, "60", "1", 0, "", true, 0, "")

	pdf.Ln(-1)

	pdf.SetFont("Times", "", 10)
	pdf.SetFillColor(255, 255, 255)

	pdf.CellFormat(40, 10, "text", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 10, "same", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 10, "text", "1", 0, "R", false, 0, "")

	pdf.Ln(-1)

	if pdf.Err() {
		log.Fatalf("failed! %s", pdf.Error())
	}

	err := pdf.OutputFileAndClose("report-example.pdf")
	if err != nil {
		log.Fatalf("error saving pdf file: %s", err)
	}
}
