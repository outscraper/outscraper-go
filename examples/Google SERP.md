# Google Search Results Scraper With Go

The library returns search results from Google based on a given search query via [Outscraper API](https://app.outscraper.com/api-docs#tag/Google-Search).

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
// Search for SERP results:
results, _ := client.GoogleSearch(map[string]string {
	"query": "buy iphone 13 TX",
	"language": "en",
	"region": "us",
})
fmt.Println(results)
```
