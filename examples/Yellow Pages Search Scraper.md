# Yellow Pages Search Scraper With Python

Returns search results from Yellow Pages. [Outscraper API](https://app.outscraper.cloud/api-docs#tag/Businesses-and-POI/paths/~1yellowpages-search/get).

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
// Search results from Yellow Pages:
results, _ := client.YellowPages(map[string]string {
	"query": "outscraper.com",
})
fmt.Println(results)
```
