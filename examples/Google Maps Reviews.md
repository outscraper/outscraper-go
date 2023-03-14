# Google Maps Reviews Scraper With Go

The library provides real-time access to the reviews from Google Maps via [Outscraper API](https://app.outscraper.com/api-docs#tag/Google-Reviews).
It allows scraping all the reviews from any place on Google Maps within seconds.

- Not limited to the official Google API limit of 5 reviews per a place
- Real time data scraping with response time less than 3s
- Sort, skip, ignore, cutoff, and other advanced parameters

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
// Get reviews of the specific place by id
results, _ := client.GoogleMapsReviews(map[string]string {
	"query": "rChIJrc9T9fpYwokRdvjYRHT8nI4",
  "reviewsLimit": "20",
  "language": "en",
})
fmt.Println(results)

// Get reviews for places found by search query
results, _ := client.GoogleMapsReviews(map[string]string {
	"query": "Memphis Seoul brooklyn usa",
  "reviewsLimit": "20",
  "limit": "20",
  "language": "en",
})
fmt.Println(results)

// Get only new reviews during last 24 hours
yesterdayTimestamp := "1657980986"
results, _ := client.GoogleMapsReviews(map[string]string {
	"query": "ChIJrc9T9fpYwokRdvjYRHT8nI4",
  "sort": "newest",
  "cutoff": yesterdayTimestamp,
  "reviewsLimit": "100",
  "language": "en",
})
fmt.Println(results)
```
