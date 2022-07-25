# Outscraper Go Library

The library provides convenient access to the [Outscraper API](https://app.outscraper.com/api-docs) from applications written in the Go language. Allows using [Outscraper's services](https://outscraper.com/services/) from your code.

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
results, _ := client.GoogleMapsSearchV2(map[string]string {
	"query": "bars ny usa",
	"limit": "10",
})
fmt.Println(results)

// Get data of the specific place by id
results, _ := client.GoogleMapsSearchV2(map[string]string {
	"query": "rChIJrc9T9fpYwokRdvjYRHT8nI4",
	"language": "en",
})
fmt.Println(results)

// Get reviews of the specific place by id
results, _ := client.GoogleMapsReviewsV3(map[string]string {
	"query": "rChIJrc9T9fpYwokRdvjYRHT8nI4",
	"reviewsLimit": "20",
	"language": "en",
})
fmt.Println(results)

// Search contacts from website
results, _ := client.EmailsAndContacts(map[string]string {
	"query": "outscraper.com",
})
fmt.Println(results)
```

[More examples](https://github.com/outscraper/outscraper-go/tree/master/examples)

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/outscraper/outscraper-go.
