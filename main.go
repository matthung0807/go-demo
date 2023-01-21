package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.Post("https://www.ptt.cc/ask/over18", map[string]string{
		"from": "/bbs/Beauty/index.html",
		"yes":  "yes",
	})

	c.OnHTML("div.title", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, a *colly.HTMLElement) {
			fmt.Println(a.Text)
		})
	})

	c.Visit("https://www.ptt.cc/bbs/Beauty/index.html")
}
