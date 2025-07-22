# Trustpilot Search Scraper With Go

Returns search results from Trustpilot. [Outscraper API](https://app.outscraper.cloud/api-docs#tag/Trustpilot/paths/~1trustpilot~1search/get).

## Installation

Go 1.10+ must be already installed.

Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):
``` sh
go mod init
```

``` go
go get -u github.com/outscraper/outscraper-go
```

[Link to the Go module page](https://pkg.go.dev/github.com/outscraper/outscraper-go)

## Initialization
```go
package main

import (
	"fmt"
	"github.com/outscraper/outscraper-go"
)

client := outscraper.Client{ApiKey: "SECRET_API_KEY"}
```
[Link to the profile page to create the API key](https://app.outscraper.com/profile)

## Usage

```go
// Get information about the search results from Trustpilot:
results, _ := client.TrustpilotSearch(map[string]string {
	"query": "real estate",
})
fmt.Println(results)
```
