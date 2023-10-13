package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator() // create pdf generator
	if err != nil {
		panic(err)
	}
	htmlfile, err := ioutil.ReadFile("./invoice_tw.html") // read html file
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))
	pdfg.Dpi.Set(600)

	err = pdfg.Create() // create pdf in memory
	if err != nil {
		panic(err)
	}

	err = pdfg.WriteFile("./invoice_tw.pdf") // create pdf file
	if err != nil {
		panic(err)
	}
}
