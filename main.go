package main

import (
	"dns_updater/client"
	googledomain "dns_updater/client/google_domain"
	"dns_updater/client/mydns"
	"dns_updater/config"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func getGlobalIp() string {
	// curl checkip.amazonaws.com
	return ""
}

func do(client client.Client, errCh chan error) {
	client.Update(errCh)
	wg.Done()
	fmt.Println("wait group was decrement")
}

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(config)
	errCh := make(chan error, len(config.Cfgs))
	defer close(errCh)
	var client client.Client
	for _, cfg := range config.Cfgs {
		switch cfg.Env {
		case "mydns":
			client = mydns.NewMyDNSClient(cfg)
		case "googleDomain":
			client = googledomain.NewGoogleDomainClient(cfg)
		default:
			fmt.Sprintln("unsupported env: %s, skipped", cfg.Env)
			continue
		}
		wg.Add(1)
		go do(client, errCh)
	}
	wg.Wait()
	select {
	case err, closed := <-errCh:
		if !closed {
			fmt.Printf("Value %s was read.\n", err.Error())
		} else {
			fmt.Println("Channel closed!")
		}
	default:
		fmt.Println("No value ready, moving on.")
	}

	fmt.Println("end.")
}
