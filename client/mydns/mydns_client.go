package mydns

import (
	ddnsclient "dns_updater/client/ddns_client"
	"dns_updater/logger"

	"dns_updater/config"
)

var myDNSURL string = "https://ipv4.mydns.jp/login.html"

type MyDNSClient struct {
	ddnsclient *ddnsclient.DDNSClient
	logger     *logger.Logger
	Name       string
}

func (m *MyDNSClient) List() []string {
	return []string{"list", "list"}
}

func (m *MyDNSClient) Add() error {
	return nil
}

func (m *MyDNSClient) Update(errCh chan<- error) {
	m.logger.Debug("update called")
	m.logger.InfoF("update mydns with %s will start", m.Name)
	resp, err := m.ddnsclient.Update()
	respS := resp.String()
	if err != nil {
		m.logger.ErrorF("update mydns with %s failed", m.Name)
		errCh <- err
	} else {
		m.logger.Info(respS)
		m.logger.InfoF("update with %s was executed successfully", m.Name)
	}
	m.logger.InfoF("update mydns with %s was executed successfully", m.Name)
}

func NewMyDNSClient(cfg *config.YamlConfig, logger *logger.Logger) *MyDNSClient {
	mydnsLogger := logger.Child("MyDNSClient")
	ddnsclient, err := ddnsclient.NewDDNSClient(myDNSURL, cfg.Timeout, cfg.MyDNS.UserName, cfg.MyDNS.Pass, mydnsLogger)

	if err != nil {
		logger.Error("create MyDNSClient failed")
		logger.Error(err.Error())
		return nil
	}

	return &MyDNSClient{ddnsclient: ddnsclient, logger: mydnsLogger, Name: cfg.MyDNS.UserName}
}
