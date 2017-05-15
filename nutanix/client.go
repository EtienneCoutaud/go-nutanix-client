package nutanix

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
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

func (c *Client) GetRequest(uri string) ([]byte, error) {
	req, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Fail to GET " + uri)
		return nil, errors.New("Return status: " + strconv.Itoa(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) PostRequest(uri string, d interface{}) ([]byte, error) {
	if d == nil {
		return nil, errors.New("POST Request, data are nil")
	}
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	resdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Fail to POST " + uri)
		return nil, errors.New("Return status: " + strconv.Itoa(resp.StatusCode))
	}
	return resdata, err
}
