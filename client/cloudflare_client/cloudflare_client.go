package cloudflare

import (
	"bytes"
	globalIPClient "dns_updater/client/my_global_ip"
	httpClient "dns_updater/client/my_http_client"
	"dns_updater/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dns_updater/config"
)

var cloudFlareURL string = "https://api.cloudflare.com/client/v4/zones"

type CloudFlareClient struct {
	httpClient *httpClient.MyHttpClient
	req        *http.Request
	logger     *logger.Logger
	cfg        *config.CloudFlareConfig
}

func (c *CloudFlareClient) List() []string {
	return []string{"list", "list"}
}

func (c *CloudFlareClient) Add() error {
	return nil
}

func (c *CloudFlareClient) SetParam(key string, value string) {
	q := c.req.URL.Query()
	q.Add(key, value)
	c.req.URL.RawQuery = q.Encode()
}

func (c *CloudFlareClient) SetHeader(req *http.Request, headerMap map[string]string) {

	if len(headerMap) == 0 {
		c.logger.Debug("header length is 0, skip set header")
		return
	}

	for key, val := range headerMap {
		req.Header.Set(key, val)
	}

}

func (c *CloudFlareClient) makeRequest(method string, url string, headerMap map[string]string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Print("create request failed")
		return nil, err
	}

	headerMap["User-Agent"] = "dns_udater/1.0.0"
	headerMap["Content-Type"] = "application/json"
	headerMap["Authorization"] = fmt.Sprintf("Bearer %", c.cfg.Token)

	c.SetHeader(req, headerMap)

	return req, nil
}

func (c *CloudFlareClient) do(req *http.Request, body io.Reader) error {
	resp, err := c.httpClient.Get(req)
	if err != nil {
		return err
	}
	fmt.Print(string(resp.String()))
	return nil
}

func (c *CloudFlareClient) makeBody(skipVerify bool) (io.Reader, error) {
	gipClient := globalIPClient.NewGlobalIPClient(5, skipVerify, c.logger)
	gip, err := gipClient.Get()
	if err != nil {
		return nil, err
	}

	data := map[string]string{}
	data["content"] = *gip
	data["name"] = c.cfg.Name
	data["proxied"] = "false"
	data["type"] = c.cfg.Type
	data["comment"] = c.cfg.Comment

	dataJson, err := json.Marshal(data)
	if err != nil {
		c.logger.Error("json marshal failed")
		return nil, err
	}

	return bytes.NewBufferString(string(dataJson)), nil
}

func (c *CloudFlareClient) Update(errCh chan<- error) {
	method := "PUT"
	url := cloudFlareURL
	body, err := c.makeBody(true)
	if err != nil {
		errCh <- err
		return
	}
	headerMap := map[string]string{}

	req, err := c.makeRequest(method, url, headerMap, body)
	if err != nil {
		errCh <- err
		return
	}

	if err := c.do(req, body); err != nil {
		errCh <- err
		return
	}

	return
}

func NewCloudFlareClient(cfg *config.YamlConfig, logger *logger.Logger) *CloudFlareClient {
	cloudflareLogger := logger.Child("CloudFlareClient")
	client := httpClient.NewMyHttpClient(cfg.Timeout, cfg.SkipVefiry, logger)

	return &CloudFlareClient{httpClient: client, logger: cloudflareLogger, cfg: &cfg.CloudFlare}
}
