package nutanix

import (
	"bytes"
	"crypto/tls"
	b64 "encoding/base64"
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
	Username   string
	Password   string
	BaseURL    url.URL
	HTTPClient *http.Client
}

// New create a new nutanix client
func New(uri string, username string, password string) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url, err := url.Parse(uri)
	if err != nil {
		log.Fatal("Failed to parse uri when generate client :", uri)
	}
	return &Client{
		Username:   username,
		Password:   password,
		BaseURL:    *url,
		HTTPClient: &http.Client{Transport: tr},
	}, nil
}

func (c *Client) NewRequest(method string, uri string, body io.Reader) (*http.Request, error) {
	url := c.BaseURL
	url.Path = path.Join(url.Path, uri)
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(c.Username+":"+c.Password)))
	req.Header.Set("Content-Type", "application/json")
	log.Println("url:", url.String())
	if body != nil {
		log.Println("body:", body.(*bytes.Buffer).String())
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) GetRequest(uri string) ([]byte, error) {
	req, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
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
	resp, err := c.HTTPClient.Do(req)
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
	return resdata, nil
}

func (c *Client) DeleteRequest(uri string) ([]byte, error) {
	req, err := c.NewRequest("DELETE", uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Fail to DELETE " + uri)
		return nil, errors.New("Return status: " + strconv.Itoa(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) PutRequest(uri string, d interface{}) ([]byte, error) {
	if d == nil {
		return nil, errors.New("PUT Request, data are nil")
	}
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("PUT", uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	resdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Fail to PUT " + uri)
		return nil, errors.New("Return status: " + strconv.Itoa(resp.StatusCode))
	}
	return resdata, nil
}
