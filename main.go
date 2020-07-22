package main

import (
	"flag"
	"fmt"
	"github.com/detectify/dtfycli/action"
	"github.com/detectify/dtfycli/api"
	"github.com/detectify/dtfycli/models"
	"net/http"
	"os"
)

func main() {

	// Parse the CLI flags
	act := "help"
	target := ""
	token := ""
	severity := ""
	printAsJSON := false
	heuristicGroup := false
	queryDaysBack := 90
	flag.StringVar(&token, "x", token, "The API key to be used (can be set using env DTFY_API_KEY)")
	flag.StringVar(&act, "a", act, "The action to execute")
	flag.StringVar(&target, "t", target, "The target hostname to query")
	flag.StringVar(&severity, "s", severity, "The severity of findings to query (high/medium/low/info)")
	flag.BoolVar(&printAsJSON, "j", printAsJSON, "Print the results as JSON (can be used with jq or gron)")
	flag.BoolVar(&heuristicGroup, "g", heuristicGroup, "Group results based on hostname/title")
	flag.IntVar(&queryDaysBack, "d", queryDaysBack, "How many days back to query findings, use 0 for all days")
	flag.Parse()

	// Try to setup the token...
	if token == "" {
		token, _ = os.LookupEnv("DTFY_API_KEY")
	}
	if token == "" {
		act = "help"
	}

	// Figure out which severity level to be used when querying findings
	s := models.SeverityAny
	switch severity {
	case "h", "high", "c", "critical":
		s = models.SeverityHigh
	case "m", "medium":
		s = models.SeverityMedium
	case "l", "low":
		s = models.SeverityLow
	case "i", "info", "information":
		s = models.SeverityInformation
	case "", "a", "any":
		s = models.SeverityAny
	default:
		os.Stderr.WriteString(fmt.Sprintf("Unrecognized severity %s, using any\n", severity))
	}

	// Pass on a stateful context containing extra configuration
	ctx := action.Context{
		PrintAsJSON:       printAsJSON,
		GroupByHeuristics: heuristicGroup,
		QueryBySeverity:   s,
		QueryDaysBack:     queryDaysBack,
	}

	// If the token was set, we can safely spin up the API client
	client := api.NewClient(token, http.DefaultClient)

	// Figure out what
	switch act {
	case "d", "domain", "domains", "ls":
		action.Domains(client, target, &ctx)
	case "f", "finding", "findings":
		action.Findings(client, target, &ctx)
	default:
		status := 0
		flag.Usage()
		fmt.Println()
		if token == "" {
			fmt.Println("No API key specified!")
			fmt.Println("Visit https://developer.detectify.com/#header-authentication for more information")
			status = 1
		} else {
			if act == "help" || act == "h" {
				fmt.Println("Possible commands:")
			} else {
				fmt.Printf("Unrecognized command '%s':\n", act)
				status = 1
			}
			fmt.Println("   h, help:     Prints all commands")
			fmt.Println("   d, domains:  Lists domains configured for asset-monitoring")
			fmt.Println("   f, findings: Lists findings")
			fmt.Println()
			fmt.Printf("Recommended usage: %s -a findings\n", os.Args[0])
		}
		os.Exit(status)
	}
}
