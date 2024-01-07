// Package client defines the client methods for getting insight of BOLD challenges.
package client

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}
