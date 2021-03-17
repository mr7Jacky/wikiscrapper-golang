// This file contains all the functions that will be used for scraping from wikipedia
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
	link = "https://en.wikipedia.org/w/index.php?search=" +
		searchKey +
		"&title=Special:Search&profile=advanced&fulltext=1&ns0=1"
	return link
}

// wikiPageScrape will find the most relevant link in searching result
func wikiPageScrape(searchLink string, target string, numLink int) []string {
	// Create list to store the scores
	var scores []float64
	// A list to store all the links
	// A list to save return links
	var allLink, retLink []string
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
		if strings.HasPrefix(link, "#") {
			return
		}
		// Calculate the score
		score := modifySigmoid(len(link) - LCS(strings.ToLower(target), strings.ToLower(link)))
		// Obtain the full URL
		absoluteURL := e.Request.AbsoluteURL(link)
		// Print the score and link
		fmt.Printf("[INFO][Score: %.10f] Link: %s\n", score, link)
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
	fmt.Printf("[INFO] Obtain total %d link\n", len(allLink))
	// Calculate the index of link with the max score
	// Another method to use: maxScoreIndex := maxArg(scores)
	// Obtain the top n highest score
	topIndex := topNIndex(numLink, scores)
	retLink = make([]string, numLink)
	for i, idx := range topIndex {
		retLink[i] = allLink[idx]
	}
	// Return the link with the max score
	return retLink
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
	// Find the first point match <p>
	firstParagraph := doc.Find("p")
	// iterates over a Selection object
	// executing a function for each matched element
	// It returns the current Selection object.
	firstParagraph.Each(func(i int, selection *goquery.Selection) {
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
	// Print the title, Check the info
	fmt.Printf("Introduction of %s:\n\n", pageTitle)
	fmt.Println(paragraphs[0])
	// Obtain the content of paragraph
	content := getContents(paragraphs, 0, len(paragraphs))

	return paragraphs[0], pageTitle, content
}

// createWikiContents is used to create Article list from given links
func createWikiContents(links []string) []Article {
	// Initialize
	var articles []Article
	var description, title, contents string
	articles = make([]Article, len(links))
	// Loop all the link and create a short article overview
	for i, link := range links {
		description, title, contents = wikiIntroScrape(link)
		articles[i] = Article{
			Author:      "WIKIPEDIA",
			Title:       title,
			Description: description,
			URL:         link,
			Content:     contents,
		}
	}
	return articles
}

// Search is the function used to search the keyword
func Search(key string) []Article {
	// Reformat search key with multiple key words
	key = reformatSearchKey(key)
	fmt.Println(key)
	// Generate the link for searching
	var wikiLink string
	wikiLink = wikiSearchLinkGenerator(key)
	// Obtain the most related links
	var topLinks []string
	topLinks = wikiPageScrape(wikiLink, key, 5)
	fmt.Printf("[INFO] The most relevant link is: %s\n", topLinks)
	// Create a wiki preview articles and link
	var articles []Article
	articles = createWikiContents(topLinks)
	return articles
}
