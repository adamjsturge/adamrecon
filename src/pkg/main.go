package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	domainsFlag := flag.String("domains", "", "Comma-separated list of domains to scan")
	flag.Parse()

	if *domainsFlag == "" {
		log.Fatal("Please provide domains using -domains flag")
	}

	domains := strings.Split(*domainsFlag, ",")
	fmt.Printf("Starting reconnaissance for domains: %v\n", domains)

	var allDomains []string
	allDomains = append(allDomains, domains...)

	// Run Findomain
	fmt.Println("\n[+] Running Findomain...")
	findomainResults := RunFindomain(domains)
	allDomains = append(allDomains, findomainResults...)

	// Run subdomain enumeration
	fmt.Println("\n[+] Running Subfinder...")
	subfinderResults := RunSubdomain(domains)
	allDomains = append(allDomains, subfinderResults...)

	// Remove duplicates
	allDomains = removeDuplicates(allDomains)
	fmt.Printf("\n[+] Total unique domains found: %d\n", len(allDomains))

	// Run remaining tools with all discovered domains
	fmt.Println("\n[+] Running HTTPx...")
	RunHttpx(allDomains)

	fmt.Println("\n[+] Running Katana crawler...")
	RunKatana(allDomains)

	// // // Run Nuclei scanner
	// fmt.Println("\n[+] Running Nuclei scanner...")
	// RunNuclei(allDomains)

	fmt.Println("\n[+] Running Subjack...")
	RunSubjack(allDomains)

	fmt.Println("\n[+] Reconnaissance completed!")
}

func removeDuplicates(domains []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, domain := range domains {
		if !seen[domain] {
			seen[domain] = true
			result = append(result, domain)
		}
	}
	return result
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
