package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RunFindomain(domains []string) []string {
	var results []string

	// Check if findomain is installed
	_, err := exec.LookPath("findomain")
	if err != nil {
		log.Println("Findomain is not installed, skipping Findomain scan")
		log.Println("Install from: https://github.com/Findomain/Findomain")
		return domains
	}

	for _, domain := range domains {
		fmt.Printf("Running Findomain scan for: %s\n", domain)

		cmd := exec.Command("findomain", "-t", domain)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error running Findomain for %s: %v\n", domain, err)
			continue
		}

		// Add found subdomains to results
		for _, result := range strings.Split(string(output), "\n") {
			if subdomain := strings.TrimSpace(result); subdomain != "" {
				results = append(results, subdomain)
			}
		}
	}

	return results
}
