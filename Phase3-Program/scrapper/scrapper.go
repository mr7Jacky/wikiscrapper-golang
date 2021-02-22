package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func scp() {
	fmt.Println("Scrapper")
	c := colly.NewCollector(
		colly.AllowedDomains("go-colly.org"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		if strings.HasPrefix(link, "/") {
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			link = strings.Join([]string{"https://", "go-colly.org", link}, "")
		}
		//fmt.Printf("%q\n",strings.Split(link, "/"))
		//if len(strings.Split(link, "/")) >
		if err := c.Visit(e.Request.AbsoluteURL(link)); err != nil {
			//fmt.Println("Error")
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL.String())
	})

	if err := c.Visit("http://go-colly.org/"); err != nil {
		fmt.Println("Error")
	}
}

func main() {
	t := NewTree("Https://", "idex")
	t.insert("https://www.google.com/", "Google")
	t.insert("https://www.google.com/Doc/", "Google-doc")
	t.insert("https://www.google.com/Car/", "Google-car")
	t.insert("https://www.google.com/News/", "Google-news")
	t.insert("https://www.google.com/Doc/News/", "Google-doc-news")
	t.insert("https://www.baidu.com/", "Baidu")
	t.PrintTree()
}
