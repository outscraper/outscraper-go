# Youtube Comments With Python

Returns comments from YouTube videos.[Outscraper API](https://app.outscraper.cloud/api-docs#tag/Reviews-and-Comments/paths/~1youtube-comments/get).

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
// Get information about the comments from YouTube videos:
results, _ := client.YoutubeComments(map[string]string {
	"query": "https://www.youtube.com/watch?v=ph5pHgklaZ0",
})
fmt.Println(results)
```
