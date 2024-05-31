package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haccer/subjack/subjack"
)

func RunSubjack() {
	var fingerprints []subjack.Fingerprints
	config, _ := ioutil.ReadFile("custom_fingerprints.json")
	json.Unmarshal(config, &fingerprints)

	subdomain := "dead.cody.su"
	/* Use subjack's advanced detection to identify
	if the subdomain is able to be taken over. */
	service := subjack.Identify(subdomain, false, false, 10, fingerprints)

	if service != "" {
		service = strings.ToLower(service)
		fmt.Printf("%s is pointing to a vulnerable %s service.\n", subdomain, service)
	}
}
