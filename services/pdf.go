package services

import (
	"fmt"
	"log"
	"manajemen-work-order/models"
	"os"
	"time"

	"github.com/signintech/gopdf"
)

func RemoveFile(path string) {
	os.Remove(path)
}

func CreatePDFofPPP(ppp models.PPP, entity models.Entity, toWhom string) (string, error) {

	//TOBE

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("arial", "assets/ttf/arial.ttf")
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	pdf.Image("assets/img/logo.png", 200, 50, nil) //print image
	err = pdf.SetFont("arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
	pdf.SetX(0) //move current location
	pdf.SetY(0)

	path := fmt.Sprintf("archive/ppp/%s%s.pdf", entity.Fullname, time.Now())

	err = pdf.WritePdf(path)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return path, nil
}
