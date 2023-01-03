package googledomain

import (
	ddnsclient "dns_updater/client/ddns_client"
	"dns_updater/config"
	"dns_updater/logger"
	"errors"
	"regexp"
)

var (
	googleDomainURL string = "https://domains.google.com/nic/update"
	okResp                 = regexp.MustCompile(`^(good|nochg)\s[0-9\.]*$`)
)

type GoogleDomainClient struct {
	ddnsclient *ddnsclient.DDNSClient
	logger     *logger.Logger
	Name       string
}

func (m *GoogleDomainClient) List() []string {
	return []string{"list", "list"}
}

func (m *GoogleDomainClient) Add() error {
	return nil
}

func (m *GoogleDomainClient) Update(errCh chan<- error) {
	m.logger.Debug("update called")
	m.logger.InfoF("update google domain with %s will start", m.Name)
	resp, err := m.ddnsclient.Update()
	if respS := resp.String(); !okResp.MatchString(resp.String()) {
		m.logger.ErrorF("update with %s failed", m.Name)
		err = errors.New(respS)
		errCh <- err
	} else if err != nil {
		m.logger.ErrorF("update with %s failed", m.Name)
		errCh <- err
	} else {
		m.logger.InfoF("update with %s was executed successfully", m.Name)
	}

}

func NewGoogleDomainClient(cfg *config.YamlConfig, logger *logger.Logger) *GoogleDomainClient {
	googleDomainLogger := logger.Child("GoogleDomainClient")
	ddnsclient, err := ddnsclient.NewDDNSClient(googleDomainURL, cfg.Timeout, cfg.GoogleDomain.UserName, cfg.GoogleDomain.Pass, googleDomainLogger)

	if err != nil {
		logger.Error("create GoogleDomainClient failed")
		logger.Error(err.Error())
		return nil
	}
	ddnsclient.SetParam("hostname", cfg.GoogleDomain.Name)
	return &GoogleDomainClient{ddnsclient: ddnsclient, logger: googleDomainLogger, Name: cfg.GoogleDomain.Name}
}
