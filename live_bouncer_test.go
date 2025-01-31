package csbouncer_test

import (
	"fmt"
	"log"

	csbouncer "github.com/blind-oracle/go-cs-bouncer"
)

func ExampleLiveBouncer() {
	bouncer := &csbouncer.LiveBouncer{
		APIKey: "ebd4db481d51525fd0df924a69193921",
		APIUrl: "http://localhost:8080/",
	}

	if err := bouncer.Init(); err != nil {
		log.Fatalf(err.Error())
	}

	response, err := bouncer.Get([]string{}, "")
	if err != nil {
		log.Fatalf("unable to get decisions: '%s'", err)
	}

	if len(*response) == 0 {
		log.Printf("no decision")
	}

	for _, decision := range *response {
		fmt.Printf("decisions: IP: %s | Scenario: %s | Duration: %s | Scope : %v\n", *decision.Value, *decision.Scenario, *decision.Duration, *decision.Scope)
	}
}
func ExampleLiveBouncer_Config() {
	bouncer := &csbouncer.LiveBouncer{}

	err := bouncer.Config("./config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	if err = bouncer.Init(); err != nil {
		log.Fatalf(err.Error())
	}

	response, err := bouncer.Get([]string{}, "")
	if err != nil {
		log.Fatalf("unable to get decision: '%s'", err)
	}

	if len(*response) == 0 {
		log.Printf("no decision")
	}

	for _, decision := range *response {
		fmt.Printf("decisions: IP: %s | Scenario: %s | Duration: %s | Scope : %v\n", *decision.Value, *decision.Scenario, *decision.Duration, *decision.Scope)
	}
}
