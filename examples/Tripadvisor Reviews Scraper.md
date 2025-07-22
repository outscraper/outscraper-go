# Tripadvisor Reviews Scraper With Go

Returns reviews from Tripadvisor businesses.
In case no reviews were found by your search criteria, your search request will consume the usage of one review. [Outscraper API](https://app.outscraper.cloud/api-docs#tag/Reviews-and-Comments/paths/~1tripadvisor-reviews/get).

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
// Get information about business:
results, _ := client.TripadvisorReviews(map[string]string {
	"query": "https://www.tripadvisor.com Restaurant_Review-g187147-d12947099-Reviews-Mayfair_Garden-Paris_Ile_de_France.html",
})
fmt.Println(results)
```
