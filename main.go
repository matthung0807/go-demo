package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("h1.title", func(e *colly.HTMLElement) {
		s := e.Text // the text of &lt;h1 class="title"/&gt;
		fmt.Println(s)
	})

	c.Visit("https://matthung0807.blogspot.com/")
}
