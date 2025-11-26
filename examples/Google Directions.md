# Google Directions With Go

Returns directions between two points from Google Maps.[Outscraper API](https://app.outscraper.cloud/api-docs#tag/Google/paths/~1maps~1directions/get).

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
# Returns directions:
results, _ := client.GoogleMapsDirections(map[string]interface{}{
		"origin": [][]string{
			{"29.696596", "76.994928"},
			{"30.715966244353", "76.8053887016268"},
		},
		"destination": [][]string{
			{"29.696596", "76.994928"},
			{"30.723065", "76.770169"},
		},
	})
fmt.Println(results)
```
