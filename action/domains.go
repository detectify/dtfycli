package action

import (
	"encoding/json"
	"fmt"
	"github.com/detectify/dtfycli/api"
	"os"
)

func Domains(client *api.Client, target string, ctx *Context) {

	// This call get all domain names (and associated metadata) from the Detectify API
	res, err := client.GetDomains()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Here we iterate over all domain objects
	status := 0
	for _, entry := range res {

		// The Context-struct contain a flag defining whether or not to print the results as JSON.
		// Any JSON result can be used in conjunction with tools such as jq or gron.
		output := ""
		if ctx.PrintAsJSON {
			bytes, err := json.Marshal(entry)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
				status = 1
				continue
			}
			output = string(bytes)
		} else {

			// The standard formatting. The UUID is the primary reference used throughout the API.
			output = fmt.Sprintf("%s\t%s", entry.UUID, entry.Name)
		}

		// The domains call return all roots present in asset-inventory, here we manually filter the results
		if target == "" {
			fmt.Println(output)
		} else if target == entry.UUID || target == entry.Name {
			fmt.Println(output)
		}
	}
	os.Exit(status)
}
