package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

func getGlobalIp() string {
	// curl checkip.amazonaws.com
	return ""
}

func main() {
	ctx := context.Background()
	credentialJson, err := os.Open("credentials.json")
	if err != nil {
		fmt.Println("credential.json not found")
		log.Fatal(err)
	}
	var credentialContent []byte
	credentialContent, err = ioutil.ReadAll(credentialJson)
	if err != nil {
		fmt.Println("read err from credential.json")
		log.Fatal(err)
	}
	credentials, err := google.CredentialsFromJSON(ctx, credentialContent)
	if err != nil {
		fmt.Println("auth err")
		log.Fatal(err)
	}

	dnsService, err := dns.NewService(ctx, option.WithCredentialsJSON(credentials.JSON))
	if err != nil {
		log.Fatal(err)
	}

	rrService := dns.NewResourceRecordSetsService(dnsService)
	resp, err := rrService.Get("project-id", "zone-name", "domain", "A").Context(ctx).Do()
	if err != nil {
		fmt.Println("can not get dns info")
		log.Fatal(err)
	}
	var pResp *dns.ResourceRecordSet

	resp.Name = "test.example.com"
	pResp, err = rrService.Patch("project-id", "zone-name", "test.example.com", "A", resp).Context(ctx).Do()
	if err != nil {
		fmt.Println("can not update dns info")
		log.Fatal(err)
	}
	fmt.Println(pResp.MarshalJSON())
}
