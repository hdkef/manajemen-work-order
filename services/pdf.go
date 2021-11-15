package services

import (
	"fmt"
	"manajemen-work-order/models"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func RemoveFile(path string) {
	os.Remove(path)
}

func createHeader(pdf *gofpdf.Fpdf) {
	pdf.ImageOptions("assets/img/ppsdm.png", 10, 6, 42, 42, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	pdf.SetFont("times", "", 8)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(0, 5, "KEMENTERIAN ENERGI DAN SUMBER DAYA MINERAL\n REPUBLIK INDONESIA", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	pdf.SetFont("times", "B", 8)
	pdf.MultiCell(0, 5, "BADAN PENGEMBANGAN SUMBER DAYA MANUSIA\nENERGI DAN SUMBER DAYA MINERAL", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	pdf.SetFont("times", "B", 10)
	pdf.MultiCell(0, 5, "PUSAT PENGEMBANGAN SUMBER DAYA MANUSIA\nMINYAK DAN GAS BUMI", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	pdf.SetFont("times", "", 8)
	pdf.MultiCell(0, 10, "JALAN SOROGO 1 CEPU, BLORA-JAWA TENGAH", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	pdf.SetFont("times", "", 8)
	pdf.MultiCell(0, 2, "TELEPON (0296)421888 FAKSIMILE: (0296)421891 http://www.ppsdmmigas.esdm.go.id E-mail: info.migas@esdm.go.id", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	w, _ := pdf.GetPageSize()
	pdf.Line(10, 55, w-10, 55)
}

func createNotaDinas(pdf *gofpdf.Fpdf, nota string) {
	pdf.SetFont("times", "", 12)
	pdf.MultiCell(0, 10, "", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	pdf.MultiCell(0, 5, "NOTA DINAS", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	pdf.MultiCell(0, 5, nota, gofpdf.BorderNone, gofpdf.AlignCenter, false)
}

func createOpening(pdf *gofpdf.Fpdf, towhom string, from string, sifat string, perihal string) {
	pdf.SetFont("times", "", 10)
	pdf.MultiCell(0, 10, "", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	pdf.MultiCell(0, 7, fmt.Sprintf("Yang terhormat\t: %s", towhom), gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.MultiCell(0, 7, fmt.Sprintf("Dari\t: %s", from), gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.MultiCell(0, 7, fmt.Sprintf("Sifat\t: %s", sifat), gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.MultiCell(0, 7, fmt.Sprintf("Perihal\t: %s", perihal), gofpdf.BorderNone, gofpdf.AlignLeft, false)
}

func createBody(pdf *gofpdf.Fpdf, pekerjaan []string) {
	pdf.SetFont("times", "", 12)
	pdf.MultiCell(0, 20, "", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	pdf.MultiCell(0, 5, "Mohon untuk dilakukan pekerjaan sebagai berikut: ", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	pdf.MultiCell(0, 10, "", gofpdf.BorderNone, gofpdf.AlignCenter, false)
	for _, v := range pekerjaan {
		pdf.MultiCell(0, 5, v, gofpdf.BorderNone, gofpdf.AlignCenter, false)
	}
}

func createSignature(pdf *gofpdf.Fpdf, fromName string, signature string) {
	w, _ := pdf.GetPageSize()
	pdf.ImageOptions(signature, w-45, 200, 25, 25, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")
	pdf.SetFont("times", "", 12)
	pdf.Text(w-65, 235, fromName)
}

func CreatePDFofPPP(ppp models.PPP, entity models.Entity, toWhom string) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	createHeader(pdf)
	createNotaDinas(pdf, ppp.Nota)
	createOpening(pdf, toWhom, entity.Role, ppp.Sifat, ppp.Perihal)
	createBody(pdf, []string{ppp.Pekerjaan})
	createSignature(pdf, entity.Fullname, entity.Signature)
	w, _ := pdf.GetPageSize()

	date := time.Now()
	pdf.SetFont("arial", "", 8)
	pdf.Text(w-50, 80, date.Format("Monday 02 January 2006"))

	path := fmt.Sprintf("archive/ppp/%s%s.pdf", entity.Fullname, time.Now())

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return "", err
	}
	return path, nil
}
