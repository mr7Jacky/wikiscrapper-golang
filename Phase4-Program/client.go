// This file contains all the functions for client
package main

import (
	"net/http"
)

// FetchEverything is used for obtain results from scrapper
func (c *Client) FetchEverything(query string) (*Results, error) {
	articleList := Search(query)
	res := &Results{
		TotalResults: len(articleList),
		Articles:     articleList,
	}
	return res, nil
}

// NewClient is used to Creat a new Client
func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient}
}
