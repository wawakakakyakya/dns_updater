package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type MyHttpClient struct {
	url    string
	client *http.Client
}

func (c *MyHttpClient) Get(req *http.Request) (*bytes.Buffer, error) {
	fmt.Println("MyHttpClient.Get called")
	fmt.Printf("MyHttpClient http request: %+v\n", req)
	var buf bytes.Buffer

	resp, err := c.client.Do(req)
	fmt.Printf("MyHttpClient request to %s was executed\n", req.URL)
	if err != nil {
		fmt.Println("MyHttpClient Get failed")
		return nil, err
	}
	defer resp.Body.Close()
	io.Copy(&buf, resp.Body)
	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return &buf, err
	} else {
		fmt.Println("MyHttpClient api was called successfully")
	}
	res, err := ioutil.ReadAll(&buf)
	if err != nil {
		fmt.Println(err.Error())
		return &buf, err
	}
	fmt.Printf("http body: \n%+v\n", string(res))
	return &buf, nil
}

func NewHttpClient(url string, timeout int) *MyHttpClient {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	return &MyHttpClient{url: url, client: client}
}
