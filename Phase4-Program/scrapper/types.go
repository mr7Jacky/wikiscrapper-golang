package main

import "net/http"

// Article represents the article found on wikipedia
type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Content     string `json:"content"`
}

// Results stores the top ranked relevant articles from the search results
type Results struct {
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Client struct {
	http     *http.Client
	PageSize int
}
