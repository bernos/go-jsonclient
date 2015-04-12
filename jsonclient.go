package jsonclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"io"
)

type Client struct {
	*http.Client
}

type Response struct {
	*http.Response	
}

func (r *Response) To(v interface{}) error {
	defer r.Body.Close()

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		return err
	} else {
		return json.Unmarshal(body, &v)
	}
}

func (c *Client) Get(url string) (*Response, error) {
	return c.execute("GET", url, nil)
}

func (c *Client) Post(url string, body interface{}) (*Response, error) {
	return c.execute("POST", url, body)
}

func (c *Client) Put(url string, body interface{}) (*Response, error) {
	return c.execute("PUT", url, body)
}

func (c *Client) Delete(url string) (*Response, error) {
	return c.execute("DELETE", url, nil)	
}

func NewJsonClient(c *http.Client) (*Client, error) {
	if c == nil {
		c = new(http.Client)
	}

	return &Client{c}, nil
}

func (c *Client) execute(method string, url string, body interface{}) (*Response, error) {

	var reader io.Reader

	if body != nil {
		if data, err := json.Marshal(body); err != nil {
			return nil, err
		} else {
			reader = bytes.NewReader(data)
		}
	}

	if req, err := http.NewRequest(method, url, reader); err != nil {
		return nil, err		
	} else {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")

		if res, err := c.Do(req); err != nil {
			return nil, err
		} else {
			return &Response{res}, nil
		}
	}
}