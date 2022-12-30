package mydns

import (
	ddnsclient "dns_updater/client/ddns_client"

	"dns_updater/config"
	"fmt"
)

var myDNSURL string = "https://ipv4.mydns.jp/login.html"

type MyDNSClient struct {
	ddnsclient *ddnsclient.DDNSClient
}

func (m *MyDNSClient) List() []string {
	return []string{"list", "list"}
}

func (m *MyDNSClient) Add() error {
	return nil
}

func (m *MyDNSClient) Update(errCh chan<- error) {
	fmt.Println("MyDNSClient.Update called")

	_, err := m.ddnsclient.Update()
	if err != nil {
		fmt.Sprintln("MyDNSClient.Update failed")
		errCh <- err
		return
	}
	return
}

func NewMyDNSClient(cfg config.YamlConfig) *MyDNSClient {
	ddnsclient, err := ddnsclient.NewDDNSClient(myDNSURL, cfg.Timeout, cfg.MyDNS.UserName, cfg.MyDNS.Pass)
	if err != nil {
		fmt.Println("create MyDNSClient failed")
		fmt.Println(err.Error())
		return nil
	}
	return &MyDNSClient{ddnsclient: ddnsclient}
}
