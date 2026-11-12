package scraper

import (
	"log"
	"rsc.io/pdf"
)

func read_pdf() {
	r, err := pdf.Open("catalog.pdf")
	if err != nil {
		log.Fatal(err)
	}
	page := r.Page(1)
	text := page.Content() // or use .Text or .V()
	_ = text // do something with text or print it
}
