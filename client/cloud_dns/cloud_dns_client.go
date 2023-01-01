package clouddns

import (
	"context"
	myglobalip "dns_updater/client/my_global_ip"
	"dns_updater/config"
	"dns_updater/logger"

	"github.com/wawakakakyakya/configloader/file"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

type CloudDNSClient struct {
	logger         *logger.Logger
	Name           *string
	globalIPClient *myglobalip.GlobalIPClient
	cfg            *config.CloudDNS
}

func NewCloudDNSClient(cfg *config.YamlConfig, logger *logger.Logger) *CloudDNSClient {
	cloudDNSLogger := logger.Child("CloudDNSClient")
	globalIPClient := myglobalip.NewGlobalIPClient(cfg.Timeout, cloudDNSLogger)
	return &CloudDNSClient{logger: cloudDNSLogger, Name: &cfg.CloudDNS.Name, globalIPClient: globalIPClient, cfg: &cfg.CloudDNS}
}

func (c *CloudDNSClient) List() []string {
	return []string{"list", "list"}
}

func (c *CloudDNSClient) Update(errCh chan<- error) {
	c.logger.InfoF("update clouddns with %s started", c.Name)

	ctx := context.Background()
	credentialJson, err := file.ReadAll(c.cfg.Credential)
	if err != nil {
		c.logger.Error("read credential failed")
		errCh <- err
		return
	}
	credentials, err := google.CredentialsFromJSON(ctx, credentialJson)
	if err != nil {
		c.logger.Error("auth error")
		errCh <- err
		return
	}

	dnsService, err := dns.NewService(ctx, option.WithCredentialsJSON(credentials.JSON))
	if err != nil {
		c.logger.Error("create service failed")
		errCh <- err
		return
	}
	globalIP, err := c.globalIPClient.Get()
	if err != nil {
		errCh <- err
		return
	}
	c.logger.Info(*globalIP)
	rrService := dns.NewResourceRecordSetsService(dnsService)
	resp, err := rrService.Get(c.cfg.ProjectID, c.cfg.ZoneName, c.cfg.Name, c.cfg.RecordType).Context(ctx).Do()
	if err != nil {
		c.logger.Error("get zone info failed")
		errCh <- err
		return
	}
	var pResp *dns.ResourceRecordSet

	resp.Name = c.cfg.Name
	pResp, err = rrService.Patch(c.cfg.ProjectID, c.cfg.ZoneName, c.cfg.Name, c.cfg.RecordType, resp).Context(ctx).Do()
	if err != nil {
		c.logger.Error("update zone info failed")
		errCh <- err
		return
	}
	j, err := pResp.MarshalJSON()
	if err != nil {
		c.logger.Error("marshal json failed")
		errCh <- err
		return
	}
	c.logger.Debug(string(j))

}
