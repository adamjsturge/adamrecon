package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haccer/subjack/subjack"
)

func RunSubjack(domains []string) {
	var fingerprints []subjack.Fingerprints
	config, _ := ioutil.ReadFile("custom_fingerprints.json")
	json.Unmarshal(config, &fingerprints)

	for _, subdomain := range domains {
		service := subjack.Identify(subdomain, false, false, 10, fingerprints)

		if service != "" {
			service = strings.ToLower(service)
			fmt.Printf("%s is pointing to a vulnerable %s service.\n", subdomain, service)
		}
	}
}
