package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func RunSubdomain(domains []string) []string {
	var results []string
	subfinderOpts := &runner.Options{
		Threads:            10, // Thread controls the number of threads to use for active enumerations
		Timeout:            30, // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10, // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Silent:             false,
	}

	// disable timestamps in logs / configure logger
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		log.Fatalf("failed to create subfinder runner: %v", err)
	}

	output := &bytes.Buffer{}
	for _, domain := range domains {
		fmt.Printf("Enumerating subdomains for: %s\n", domain)
		if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
			log.Printf("failed to enumerate domain %s: %v", domain, err)
			continue
		}
	}

	// Add found subdomains to results
	for _, subdomain := range strings.Split(output.String(), "\n") {
		if sub := strings.TrimSpace(subdomain); sub != "" {
			results = append(results, sub)
		}
	}

	return results
}
