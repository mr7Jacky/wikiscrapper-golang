package main

import (
	// import standard libraries
	"fmt"
	"log"
	"strings"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// wikiSearchLinkGenerator will concatenate the input string with search link
func wikiSearchLinkGenerator(searchKey string) string {
	link := "https://en.wikipedia.org/wiki/Special:Search?search=" +
		searchKey +
		"&fulltext=Search+Portal+namespace&fulltext=Search&ns0=1"
	return link
}

// wikiPageScrape will find the most relevant link in searching result
func wikiPageScrape(searchLink string, target string) string {
	// Create list to store the scores
	var scores []float64
	// A list to store all the links
	var allLink []string
	// Create a new collector object
	collector := colly.NewCollector(
		// Limited to visiting domain to wikipedia
		colly.AllowedDomains("en.wikipedia.org"),
	)
	// OnHTML registers a function
	// Function will be executed on every HTML element matched by the GoQuery Selector parameter.
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Obtain the link
		link := e.Attr("href")
		// Calculate the score
		score := modifySigmoid(len(link) - LCS(target, link))
		// Obtain the full URL
		absoluteURL := e.Request.AbsoluteURL(link)
		// Print the score and link
		fmt.Printf("[Score: %.3f] Link: %s\n", score, absoluteURL)
		// Append the score
		scores = append(scores, score)
		// Obtain the link
		allLink = append(allLink, absoluteURL)
	})
	// Visiting the search page of wikipedia
	err := collector.Visit(searchLink)
	// Error Handler
	if err != nil {
		log.Fatal(err)
	}
	// Print the total number of link obtained
	fmt.Printf("[Info] Obtain total %d link\n", len(allLink))
	// Calculate the index of link with the max score
	maxScoreIndex := maxArg(scores)
	// Return the link with the max score
	return allLink[maxScoreIndex]
}

// wikiIntroScrape scrape the introduction part of the wikipedia,
// that is all the paragraph before the Contents list
func wikiIntroScrape(wikiLink string) (description string, title string, contents string) {
	// Initialize a website scrapper with the given website link
	var doc, err = goquery.NewDocument(wikiLink)
	// Error Handler
	if err != nil {
		log.Fatal(err)
	}
	// Variable store the page title
	var pageTitle string
	// Variable store the content of the paragraphs
	var paragraphs []string
	// Find() gets the descendants of each element in the current set of matched elements,
	// filtered by a selector. It returns a new Selection object containing these matched elements.
	// We use Find() to obtain the title of the website
	pageTitle = doc.Find("h1").Contents().Text()
	// Find the end point of the introduction Part which is the contents with id toc
	IntroEnd := doc.Find("div#toc")
	// gets all the preceding siblings of each element in the Selection (IntroEndPrev)
	// filtered by a selector. It returns a new Selection object containing the matched elements.
	IntroEndPrev := IntroEnd.PrevAllFiltered("p")
	// iterates over a Selection object
	// executing a function for each matched element
	// It returns the current Selection object.
	IntroEndPrev.Each(func(i int, selection *goquery.Selection) {
		// Check if the current Selection is only contains spaces, including \t, \n <space>
		curText := strings.TrimSpace(selection.Text())
		if len(curText) == 0 {
			// If so do not add to the paragraph shown later
			return
		} else {
			// Else append to our paragraphs to show later
			paragraphs = append(paragraphs, selection.Text())
		}

	})
	// Print the title
	fmt.Printf("Introduction of %s:\n\n", pageTitle)
	// Print the content of paragraph
	content := printContents(paragraphs, 0, len(paragraphs))
	return paragraphs[len(paragraphs)-1], pageTitle, content
}

// Search is the function used to search the keyword
func Search(key string) (description string, title string, contents string, wikiLink string) {
	wikiLink = wikiSearchLinkGenerator(key)
	wikiLink = wikiPageScrape(wikiLink, key)
	fmt.Printf("[INFO] The most relevant link is: %s\n", wikiLink)
	description, title, contents = wikiIntroScrape(wikiLink)
	return description, title, contents, wikiLink
}
