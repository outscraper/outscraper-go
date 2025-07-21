# Similarweb With Python

Returns website analytics data including traffic, rankings, audience insights, and competitive intelligence from SimilarWeb [Outscraper API](https://app.outscraper.cloud/api-docs#tag/Domain-Related/paths/~1similarweb/get).

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
// Get data from Similarweb businesses:
results, _ := client.Similarweb(map[string]string {
	"query": "apple.com",
})
fmt.Println(results)
```
