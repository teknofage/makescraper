package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	
	
	"github.com/jinzhu/gorm"


	"github.com/gocolly/colly"
)

// Stores the db value globally so it can be called
var dB *gorm.DB

type article struct {
	gorm.Model
	Body string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		// colly.AllowedDomains("reddit.com"),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("tbody > tr > td > table > tbody", func(e *colly.HTMLElement) {
                fmt.Println("We did it!")
				e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
					title := h.ChildText("td.title > a") 
					fmt.Println(title)
				})
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")
}
