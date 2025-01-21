package request

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Token string
	Base  string
}

func (c *Client) Request(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, body)
	request.Header.Set("X-API-KEY", c.Token)
	return client.Do(request)
}

func (c *Client) Get(url string) string {
	endpoint := c.Base + url
	response, err := c.Request("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	responseBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return string(responseBody)
}
