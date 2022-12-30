package ddnsclient

import (
	"bytes"
	client "dns_updater/client/my_http_client"
	"fmt"
	"net/http"
)

type DDNSClient struct {
	req    *http.Request
	client *client.MyHttpClient
}

func (m *DDNSClient) List() []string {
	return []string{"list", "list"}
}

func (m *DDNSClient) Add() error {
	return nil
}

func (m *DDNSClient) SetParam(key string, value string) {
	q := m.req.URL.Query()
	q.Add(key, value)
	m.req.URL.RawQuery = q.Encode()
}

func (m *DDNSClient) Update() (*bytes.Buffer, error) {
	fmt.Println("DDNSClient.Update called")
	body, err := m.client.Get(m.req)
	if err != nil {
		fmt.Sprintln("DDNSClient.Update failed")
		fmt.Println(err.Error())
		return body, err
	}
	return body, err
}

func NewDDNSClient(url string, timeout int, userName, password string) (*DDNSClient, error) {
	httpClient := client.NewHttpClient(url, timeout)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("create DDNSClient failed")
		return nil, err
	}
	req.SetBasicAuth(userName, password)
	req.Header.Set("User-Agent", "dns_udater/1.0.0")
	return &DDNSClient{req: req, client: httpClient}, nil
}
