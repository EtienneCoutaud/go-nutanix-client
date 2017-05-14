package client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

// Client use to perform request
type Client struct {
	BaseURL url.URL
	*http.Client
}

// New create a new nutanix client
func New() (*Client, error) {
	return nil, nil
}

func (c *Client) NewRequest(method string, uri string, body io.Reader) (*http.Request, error) {
	url := c.BaseURL
	url.Path = path.Join(url.Path, uri)
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Add authent part
	// if c.key != "" {
	// 	req.Header.Add("Authorization", c.key)
	// }
	log.Println("url:", url.String())
	if body != nil {
		log.Println("body:", body.(*bytes.Buffer).String())
	}
	req.Header.Add("Content-Type", "application/json")
	return req, err
}
