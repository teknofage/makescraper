package main

import (
	"fmt"
	"encoding/json"
	// "io/ioutil"
	"os"
	"io"


	"github.com/gocolly/colly"
)


type article struct {
	Title 		string	`json:"title"`
	URL 		string	`json:"URL"`
	Score 		string 	`json:"score"`
	Comments 	string 	`json:"score"`
	Poster 		string	`json:"poster"`
}

func main() {
	var articles []article
	var pageCount int
	i := 0
	
	// Instantiate default collector
	c := colly.NewCollector()
	cssSelector := "tbody > tr:nth-child(3) > td > table > tbody"
	


	c.OnHTML(cssSelector, func(e *colly.HTMLElement) {
				e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
					var art article
					title := h.ChildText("td.title > a") 
					score := h.ChildText("td.subtext > span.score")
					if title == "More" {
						c.Visit("news.ycombinator.com/" + h.ChildAttr("td.title > a", "href"))
					} else if score != "" {
						articles[i].Score = score
						articles[i].Comments = h.ChildText("td.subtext > a:last-child")
						articles[i].Poster = h.ChildText("td.subtext > a.hnuser")
						i++
					} else if title != "" {
						
						art.Title = title
						art.URL = h.ChildAttr("td.title > a", "href")
						articles = append(articles, art)
					}
					
				})

	})
	// Before making a request print "Visiting ..."
	c.OnResponse(func(r *colly.Response) {
		pageCount++
        urlVisited := r.Request.URL
        fmt.Println(fmt.Sprintf("%d  DONE Visiting : %s", pageCount, urlVisited))
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")
	fmt.Println(articles[0])
	fmt.Println("We did it!")

	// Marshal instances of articles and conert to JSON
	articleJSON, _ := json.Marshal(articles)
	// fmt.Println(string(articleJSON))
	articleJSONString := string(articleJSON)
	// fmt.Println(articleJSONString)
	writeJSONToFile(articleJSONString)
}



func writeJSONToFile(articleJSONString string) error {
    file, err := os.Create("output.json")
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, articleJSONString)
    if err != nil {
        return err
    }
    return file.Sync()
}
