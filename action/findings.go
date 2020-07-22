package action

import (
	"encoding/json"
	"fmt"
	"github.com/detectify/dtfycli/api"
	"os"
	"time"
)

func Findings(client *api.Client, target string, ctx *Context) {
	res, err := client.GetDomains()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := 0
	for _, domain := range res {
		if len(target) == 0 {
			s := findingsPrint(client, domain.Token, ctx)
			if status == 0 && s != 0 {
				status = s
			}
			continue
		} else if domain.Name == target || domain.UUID == target {
			status = findingsPrint(client, domain.Token, ctx)
			break
		}
	}
	os.Exit(status)
}

func findingsPrint(client *api.Client, domainToken string, ctx *Context) int {

	// This demonstrates how to start querying from a given date. Pass in "nil" or simply discard the
	// &from= parameter to query over all time. If a negative value or zero was passed via ctx.QueryDaysBack,
	// then we can replicate this behavior and any finding will be returned.
	var from *time.Time
	if ctx.QueryDaysBack > 0 {
		now := time.Now()
		now.AddDate(0, 0, -1*ctx.QueryDaysBack) // Make the input number (QueryDaysBack) negative
		from = &now
	}

	// This call extracts all finding UUID's associated with the target domain token using a given severity and
	// date interval. Above we demonstrate the from-parameter.
	uuids, err := client.GetFindingUUIDs(domainToken, from, nil, ctx.QueryBySeverity)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
		return 1
	}

	// Below we iterate over each and every UUID and retrieve the finding data associated
	heuristicGroup := make(map[string]struct{})
	status := 0
	for _, uuid := range uuids {
		finding, err := client.GetFinding(domainToken, uuid)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
			status = 1
			continue
		}

		// This demonstrates a basic example of how to only return each unique hostname/finding tuple
		if ctx.GroupByHeuristics {
			key := fmt.Sprintf("%s:%s", finding.Title, finding.FoundAt)
			if _, found := heuristicGroup[key]; found {
				continue
			}
			heuristicGroup[key] = struct{}{}
		}

		// Here we simply return the data as JSON to be used in conjunction with another CLI tool, such as jq or gron
		if ctx.PrintAsJSON {
			serialized, err := json.Marshal(finding)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
				status = 1
				continue
			}
			fmt.Println(string(serialized))
		} else {

			// Normalize the timestamps and default to N/A if no time was provided. This can happen if Detectify
			// have found a finding that remains unpatched, in which case the EndTimestamp will be nullified.
			tsStart := "N/A"
			if finding.StartTimestamp.Unix() > 0 {
				tsStart = finding.StartTimestamp.String()
			}
			tsEnd := "N/A"
			if finding.EndTimestamp.Unix() > 0 {
				tsStart = finding.EndTimestamp.String()
			}

			// If no special formatting is applied, we simply print out some fields (can be used to make a .tsv-file)
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", finding.UUID, tsStart, tsEnd, finding.FoundAt, finding.Title)
		}
	}
	return status
}
