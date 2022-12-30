package client

import (
	"fmt"
	"net/http"
	"time"
)

type MyHttpClient struct {
	url string
	cli *http.Client
}

func (c *MyHttpClient) Get() error {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return err
	}

	fmt.Println("request to mydns was executed")
	resp, err := c.cli.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()
	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return err
	} else {
		fmt.Printf("mydns record was updated successfully")
	}
	return nil
}

func NewHttpClient(url string, timeout int) *MyHttpClient {
	cli := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	return &MyHttpClient{url: url, cli: cli}
}

func NewMyDNSClient(url string) *MyHttpClient {
	return &MyHttpClient{url: url}
}
