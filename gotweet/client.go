package gotweet

import "net/http"

type Client struct {
	Search *SearchService
}

func NewClient(client *http.Client) *Client {
	return &Client{
		Search: newSearchService(client),
	}
}
