# Company Website Finder With Go

Finds company websites based on business names.[Outscraper API](https://app.outscraper.cloud/api-docs#tag/Domain-Related/paths/~1company-website-finder/get).

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
results, _ := client.CompanyWebsiteFinder(map[string]string {
	"query": "Apple Inc",
})
fmt.Println(results)
```
