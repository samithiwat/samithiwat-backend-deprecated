package httpClient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GET(string) (*[]byte, error)
	POST(string, interface{}) (*[]byte, error)
	PATCH(string, interface{}) (*[]byte, error)
	PUT(string, interface{}) (*[]byte, error)
	DELETE(string) (*[]byte, error)
}

type HttpClient struct{}

func NewClient() {

}

func (c *HttpClient) GET(url string) (*[]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &body, err
}
func (c *HttpClient) POST(url string, contentType string, data []byte) (*[]byte, error) {
	res, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &body, err
}

func (c *HttpClient) PATCH(url string, contentType string, data []byte) (*[]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", contentType)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &body, err
}

func (c *HttpClient) PUT(url string, contentType string, data []byte) (*[]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", contentType)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &body, err
}

func (c *HttpClient) DELETE(url string) (*[]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &body, err
}
