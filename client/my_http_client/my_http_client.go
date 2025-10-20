package my_http_client

import (
	"bytes"
	"crypto/tls"
	"dns_updater/logger"
	"io"
	"net/http"
	"time"
)

type MyHttpClient struct {
	client *http.Client
	logger *logger.Logger
}

func (c *MyHttpClient) Get(req *http.Request) (*bytes.Buffer, error) {
	c.logger.Debug("Get called")
	c.logger.DebugF("http request: %+v\n", req)
	var buf bytes.Buffer

	resp, err := c.client.Do(req)
	c.logger.DebugF("request to %s was executed", req.URL)
	if err != nil {
		c.logger.Error("Get failed")
		return nil, err
	}
	defer resp.Body.Close()
	io.Copy(&buf, resp.Body)
	c.logger.DebugF("resp.StatusCode: %d", resp.StatusCode)
	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		c.logger.ErrorF("request to %s failed, status code: %d", req.URL, resp.StatusCode)
		c.logger.Error(buf.String())
		return &buf, err
	} else {
		c.logger.DebugF("request to %s was executed successfully", req.URL)
		c.logger.DebugF("response body: %s", buf.String())
	}

	return &buf, nil
}

func NewMyHttpClient(timeout int, skipVerify bool, logger *logger.Logger) *MyHttpClient {
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second, Transport: tp}
	httpClientLogger := logger.Child("HttpClient")
	return &MyHttpClient{client: client, logger: httpClientLogger}
}
