package my_http_client

import (
	"bytes"
	"dns_updater/logger"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type MyHttpClient struct {
	client *http.Client
	logger *logger.Logger
}

func (c *MyHttpClient) Get(req *http.Request) (*bytes.Buffer, error) {
	c.logger.Debug("Get called")
	c.logger.Debug(fmt.Sprintf("http request: %+v\n", req))
	var buf bytes.Buffer

	resp, err := c.client.Do(req)
	c.logger.Debug(fmt.Sprintf("request to %s was executed\n", req.URL))
	if err != nil {
		c.logger.Error("Get failed")
		return nil, err
	}
	defer resp.Body.Close()
	io.Copy(&buf, resp.Body)
	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		c.logger.Error(fmt.Sprintf("status code(%d) was not succeeded\n", resp.StatusCode))
		return &buf, err
	} else {
		c.logger.Debug("Mapi was called successfully")
	}
	res, err := ioutil.ReadAll(&buf)
	if err != nil {
		return &buf, err
	}
	c.logger.Debug(fmt.Sprintf("http body: \n%+s\n", string(res)))
	return &buf, nil
}

func NewMyHttpClient(timeout int, logger *logger.Logger) *MyHttpClient {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	httpClientLogger := logger.Child("HttpClient")
	return &MyHttpClient{client: client, logger: httpClientLogger}
}
