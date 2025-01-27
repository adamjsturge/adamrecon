package main

// import (
// 	"time"

// 	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
// )

// func initializeNucleiEngine() (*nuclei.NucleiEngine, error) {
// 	return nuclei.NewNucleiEngine(
// 		nuclei.WithTemplateFilters(nuclei.TemplateFilters{ProtocolTypes: "http,dns"}),
// 		nuclei.WithGlobalRateLimit(1, time.Second),
// 		nuclei.WithConcurrency(nuclei.Concurrency{
// 			TemplateConcurrency:           1,
// 			HostConcurrency:               1,
// 			HeadlessHostConcurrency:       1,
// 			HeadlessTemplateConcurrency:   1,
// 			JavascriptTemplateConcurrency: 1,
// 			TemplatePayloadConcurrency:    1,
// 			ProbeConcurrency:              1,
// 		}),
// 	)
// }

// func RunNuclei(domains []string) {
// 	ne, err := initializeNucleiEngine()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ne.Close()

// 	ne.LoadTargets(domains, false)

// 	err = ne.ExecuteWithCallback(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
