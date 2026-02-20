# Businesses Search With Go

Allows searching and retrieving business data via [Outscraper API](https://app.outscraper.com/api-docs#tag/Businesses).

The endpoint supports 3 request modes:
- **Structured JSON** request with `filters` and `fields`
- **AI plain-text** request with `query`
- **Combined mode** (`filters`/`fields` + `query`) where the server merges parameters

> In combined mode, `filters` and `fields` are merged. If plain text includes `limit`, `cursor`, or `include_total`, those values take priority.

---

## Installation

Go 1.10+ must be already installed.

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):
```sh
go mod init
```

```go
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

---

## Usage

### Search with structured filters (JSON)
```go
filters := map[string]interface{}{
	"country_code": "US",
	"states": []string{"NY"},
	"cities": []string{"New York", "Buffalo"},
	"types": []string{"restaurant", "cafe"},
	"has_website": true,
	"has_phone": true,
	"business_statuses": []string{"operational"},
}

page, err := client.BusinessesSearch(
	filters,
	50,
	false,
	"",
	[]string{"name", "phone", "website", "address", "rating", "reviews"},
	false,
	false,
	"",
	"",
)

fmt.Println(page, err)
```

### Search with AI plain text (`query`)
```go
page, err := client.BusinessesSearch(
	nil,
	10,
	false,
	"",
	nil,
	false,
	false,
	"",
	"Find cafes in New York, NY. Limit 25. Return name, address, phone, website, rating and reviews.",
)

fmt.Println(page, err)
```

### Combine JSON + plain text (merged request)
```go
filters := map[string]interface{}{
	"country_code": "US",
	"states": []string{"CA"},
	"types": []string{"restaurant"},
}

page, err := client.BusinessesSearch(
	filters,
	15,
	false,
	"",
	[]string{"name", "phone"},
	false,
	false,
	"",
	"Add cafes too. Return address and reviews. Limit 20. Include total.",
)

fmt.Println(page, err)
```

### Iterate over all results (auto-pagination)
```go
filters := map[string]interface{}{
	"country_code": "US",
	"states": []string{"NY"},
	"business_statuses": []string{"operational"},
}

all, err := client.BusinessesIterSearch(
	filters,
	100,
	[]string{"name", "phone", "address", "rating", "reviews"},
	false,
)

fmt.Println(len(all), err)
```

### Get business details by ID
```go
details, err := client.BusinessesGet(
	"YOUR_BUSINESS_ID",
	[]string{"name", "phone", "website", "address", "rating", "reviews"},
	false,
	false,
	"",
)

fmt.Println(details, err)
```
