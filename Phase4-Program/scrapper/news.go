package main

import (
	"net/http"
)

type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Content     string `json:"content"`
}

type Results struct {
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Client struct {
	http     *http.Client
	PageSize int
}

func (c *Client) FetchEverything(query string) (*Results, error) {
	description, title, contents, link := Search(query)
	a := Article{
		Author:      "WIKIPEDIA",
		Title:       title,
		Description: description,
		URL:         link,
		Content:     contents,
	}
	alist := []Article{a}
	res := &Results{
		TotalResults: 1,
		Articles:     alist,
	}
	return res, nil //json.Unmarshal(body, res)
}

func NewClient(httpClient *http.Client, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, pageSize}
}
