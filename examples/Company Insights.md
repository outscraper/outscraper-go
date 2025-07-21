# Company Insights With Python

Finds company details such as revenue, size, founding year, public status, etc. [Outscraper API](https://app.outscraper.cloud/api-docs#tag/Other-Services/paths/~1company-insights/get).

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
// Get information about the companies:
results, _ := client.CompanyInsights(map[string]string {
	"query": "outscraper.com",
})
fmt.Println(results)
```
