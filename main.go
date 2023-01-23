package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.Async = true

	c.Post("https://www.ptt.cc/ask/over18", map[string]string{
		"from": "/bbs/Beauty/index.html",
		"yes":  "yes",
	})

	var pageUrls []string

	next(c, &pageUrls, 0, 3) // max recursive call 3 times

	c.Visit("https://www.ptt.cc/bbs/Beauty/index.html")
	c.Wait()
	fmt.Println(strings.Join(pageUrls, "\n"))

}

func next(c *colly.Collector, pageUrls *[]string, num, max int) {
	c.OnHTML("div.btn-group.btn-group-paging", func(d *colly.HTMLElement) {
		num++
		if num > max {
			return
		}
		d.ForEach("a.btn.wide", func(i int, a *colly.HTMLElement) {
			if strings.Contains(a.Text, "上頁") {
				pageUrl := a.Attr("href")
				*pageUrls = append(*pageUrls, pageUrl)
				pageUrl = "https://www.ptt.cc" + pageUrl
				c.Visit(pageUrl)
			}
		})
	})
}
