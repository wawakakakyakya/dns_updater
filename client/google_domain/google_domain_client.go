package googledomain

import (
	ddnsclient "dns_updater/client/ddns_client"

	"dns_updater/config"
	"fmt"
)

var googleDomainURL string = "https://domains.google.com/nic/update"

type GoogleDomainClient struct {
	ddnsclient *ddnsclient.DDNSClient
}

func (m *GoogleDomainClient) List() []string {
	return []string{"list", "list"}
}

func (m *GoogleDomainClient) Add() error {
	return nil
}

func (m *GoogleDomainClient) Update(errCh chan<- error) {
	fmt.Println("GoogleDomainClient.Update called")

	_, err := m.ddnsclient.Update()
	if err != nil {
		fmt.Sprintln("GoogleDomainClient.Update failed")
		errCh <- err
		return
	}
	return
}

func NewGoogleDomainClient(cfg config.YamlConfig) *GoogleDomainClient {
	ddnsclient, err := ddnsclient.NewDDNSClient(googleDomainURL, cfg.Timeout, cfg.GoogleDomain.UserName, cfg.GoogleDomain.Pass)
	if err != nil {
		fmt.Println("create GoogleDomainClient failed")
		fmt.Println(err.Error())
		return nil
	}
	ddnsclient.SetParam("hostname", cfg.GoogleDomain.Domain)
	return &GoogleDomainClient{ddnsclient: ddnsclient}
}
