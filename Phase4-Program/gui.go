// This file contains all the functions to create a web-based gui based on html and css files
package main

import (
	"bytes"
	"html/template"
	"net/http"
	"net/url"
)

// Load the html file
var tpl = template.Must(template.ParseFiles("assets/index.html"))

// indexHandler is used for handling the index of website
func indexHandler(w http.ResponseWriter, _ *http.Request) {
	// write buffer to the content
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// searchHandler handle the search inquiries and send the search result back to web
func searchHandler(client *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtain the request
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Obtain the search key
		params := u.Query()
		searchQuery := params.Get("q")
		// Obtain the search results
		results, err := client.FetchEverything(searchQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Create search result object
		search := &SearchRet{
			Query:   searchQuery,
			Results: results,
		}
		// Write to the buffer
		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = buf.WriteTo(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
