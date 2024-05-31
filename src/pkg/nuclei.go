package main

import (
	"sync"

	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
)

func RunNuclei() {
	ne, err := nuclei.NewThreadSafeNucleiEngine()
	if err != nil {
		panic(err)
	}
	// setup waitgroup to handle concurrency
	wg := &sync.WaitGroup{}

	// scan 1 = run dns templates on scanme.sh
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = ne.ExecuteNucleiWithOpts([]string{"scanme.sh"}, nuclei.WithTemplateFilters(nuclei.TemplateFilters{ProtocolTypes: "http"}))
		if err != nil {
			panic(err)
		}
	}()

	// scan 2 = run http templates on honey.scanme.sh
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = ne.ExecuteNucleiWithOpts([]string{"honey.scanme.sh"}, nuclei.WithTemplateFilters(nuclei.TemplateFilters{ProtocolTypes: "dns"}))
		if err != nil {
			panic(err)
		}
	}()

	// wait for all scans to finish
	wg.Wait()
	defer ne.Close()
}
