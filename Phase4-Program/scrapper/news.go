package main

import (
	"net/http"
)

func (c *Client) FetchEverything(query string) (*Results, error) {

	aList := Search(query)
	res := &Results{
		TotalResults: 1,
		Articles:     aList,
	}
	return res, nil
}

func NewClient(httpClient *http.Client, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}
	return &Client{httpClient, pageSize}
}
