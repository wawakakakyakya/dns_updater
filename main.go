package main

import (
	"dns_updater/client"
	clouddns "dns_updater/client/cloud_dns"
	googledomain "dns_updater/client/google_domain"
	"dns_updater/client/mydns"
	"dns_updater/config"
	"dns_updater/logger"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func do(client client.Client, errCh chan error) {
	defer wg.Done()
	client.Update(errCh)
	fmt.Println("wait group was decrement")
}

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println("[ERROR] read config failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	logger := logger.NewLogger("main", &config.GlobalCfg.Log)
	errCh := make(chan error, len(config.Cfgs))
	defer close(errCh)
	var client client.Client
	for _, cfg := range config.Cfgs {
		switch cfg.Env {
		case "mydns":
			client = mydns.NewMyDNSClient(cfg, logger)
		case "googleDomain":
			client = googledomain.NewGoogleDomainClient(cfg, logger)
		case "cloudDNS":
			client = clouddns.NewCloudDNSClient(cfg, logger)
		default:
			logger.Warn(fmt.Sprintf("unsupported env: %s, skipped", cfg.Env))
			continue
		}
		wg.Add(1)
		logger.Debug("add wait group")
		go do(client, errCh)
	}

	wg.Wait()
	select {
	case err, closed := <-errCh:
		if !closed {
			logger.Info("channel closed.")
		} else {
			logger.Error(err.Error())
		}
	default:
		logger.Info("No value ready, moving on.")
	}

	logger.Info("end.")
}
