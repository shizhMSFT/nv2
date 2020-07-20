package registry

import (
	"net/http"
)

// Client is a customized registry client
type Client struct {
	base http.RoundTripper
}

// NewClient creates a new registry client
func NewClient(base http.RoundTripper, username, password string) *Client {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Client{
		base: newV2transport(base, username, password),
	}
}
