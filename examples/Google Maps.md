# Google Maps Scraper With Go

The library provides real-time access to the places from Google Maps via [Outscraper API](https://app.outscraper.com/api-docs#tag/Google-Maps).

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
// Search for businesses in specific locations:
results, _ := client.GoogleMapsSearch(map[string]string {
	"query": "restaurants brooklyn usa",
	"limit": "20",
  "language": "en",
  "region": "us",
})
fmt.Println(results)

// Get data of the specific place by id
results, _ := client.GoogleMapsSearch(map[string]string {
	"query": "ChIJrc9T9fpYwokRdvjYRHT8nI4",
  "language": "en",
})
fmt.Println(results)
```
