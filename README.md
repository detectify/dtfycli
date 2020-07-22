## Description

This is a command line utility that demonstrates how to interact with various parts of the Detectify API. We will expand this over time with more and more features. The initial release contain the ability to query domain names present in the asset-inventory offering, as well as getting asset-monitoring findings for a given domain.

## Compilation

On Linux:
* To compile this application, run the following in the project directory: `go build -o ./dtfycli`
* Mark the compiled binary as an executable: `chmod +x ./dtfycli`
* To run the application, simply type: `./dtfycli`

On Windows:
* To compile, run the follwing in the project directory: `go build -o ./dtfycli.exe`
* To run the binary, simply type: `dtfycli.exe`

## Usage

To execute any action, you must have a valid API key. For more information on how to get one [follow this link](https://developer.detectify.com/#header-authentication). The keys can be generated from the [team page in Detectify](https://detectify.com/dashboard/team). To use the key with this application, either pass it through the `-x <key>` CLI argument, or define it through the `DTFY_API_KEY`-environment variable.

This application have a few different sub-commands/actions which allows you to interact with/query various parts of the Detectify API. You can define the action via the `-a <action>` CLI argument. At the time of writing, three different actions are supported, namely: `help`, `domains` and `findings`.

To list the help section, run `dtfycli -a help`, example output:
```
Usage of ./dtfycli:
  -a string
        The action to execute (default "help")
  -d int
        How many days back to query findings, use 0 for all days (default 90)
  -g    Group results based on hostname/title
  -j    Print the results as JSON (can be used with jq or gron)
  -s string
        The severity of findings to query (high/medium/low/info)
  -t string
        The target hostname to query
  -x string
        The API key to be used (can be set using env DTFY_API_KEY)

Possible commands:
   h, help:     Prints all commands
   d, domains:  Lists domains configured for asset-monitoring
   f, findings: Lists findings

Recommended usage: ./dtfycli -a findings
```

To get a list of all domain names present in asset-inventory, run:
```
./dtfycli -a domains
8216a613-bba0-422c-839b-ee908979688f    example.com
c602f889-c0a5-4bf9-8bd9-c09e91011a52    example.net
...
```

To get a list of any findings found in asset-monitoring for example.com, run:
```
./dtfycli -a findings -t example.com
d7b1c151-dd2f-4f41-b2c8-020a3e28dbe8    2020-06-02 09:16:55 +0000 UTC   N/A     y.example.com.        NGINX Status Disclosure
c7f3cc41-ce22-424c-b2c8-d20afeef8be8    2020-06-15 11:56:24 +0000 UTC   N/A     x.example.com.        DNS-hijack, non-responding nameserver(s) in AWS Route53
```

To get a list of all high severity findings observed over the last 30 days, run:
```
./dtfycli -a findings -s high -d 30
c7f3cc41-ce22-424c-b2c8-d20afeef8be8    2020-06-15 11:56:24 +0000 UTC   N/A     x.example.com.        DNS-hijack, non-responding nameserver(s) in AWS Route53
```