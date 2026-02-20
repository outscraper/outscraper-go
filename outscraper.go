package outscraper

import (
    "bytes"
    "fmt"
    "net/url"
    "net/http"
    "encoding/json"
    "errors"
    "strings"
)

const (
    ApiUrl string = "https://api.app.outscraper.com"
)

type Client struct {
	ApiKey string
}

type Request map[string]string

type Result map[string]interface{}

type ResultArray []interface{}

func (c Client) getAPIRequest(path string, parameters map[string]string) (Result, error) {
    httpClient := &http.Client{}
    url, _ := url.Parse(ApiUrl + path)

    query := url.Query()
	if parameters != nil {
		for k, v := range parameters {
            query.Set(k, v)
		}
	}

    url.RawQuery = query.Encode()

    fmt.Println(url.String())

    req, _ := http.NewRequest("GET", url.String(), nil)
    req.Header.Add("X-API-KEY", c.ApiKey)
    req.Header.Add("Client", "Go SDK")

    response, doError := httpClient.Do(req)
	if doError != nil {
		return nil, errors.New("Cannot create request")
	}

    defer response.Body.Close()

    decoder := json.NewDecoder(response.Body)
	var result Result

	decodeError := decoder.Decode(&result)
	if decodeError != nil {
		return nil, errors.New("cannot decode")
	}

    return result, nil
}

func (c Client) GetRequestArchive(requestId string) (Result, error) {
    return c.getAPIRequest("/requests/"+ requestId, map[string]string {})
}

func (c Client) GoogleSearch(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/google-search-v3", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleSearchNews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/google-search-news", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsSearch(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/search-v2", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsSearchV3(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/search-v3", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsDirections(parameters map[string]interface{}) ([]interface{}, error) {
    q := url.Values{}

    for key, val := range parameters {
        switch v := val.(type) {
        case []string:
            for _, item := range v {
                q.Add(key, item)
            }
        case string:
            q.Add(key, v)
        }
    }

    q.Set("async", "false")

    response, err := c.getAPIRequest("/maps/directions", q)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/reviews-v3", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsPhotos(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/photos-v3", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GooglePlayReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/google-play/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) ContactsAndLeads(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/contacts-and-leads", parameters)
    if err != nil {
        return nil, err
    }

    data, ok := response["data"].([]interface{})
    if !ok {
        return nil, errors.New("unexpected response format: data is not an array")
    }

    return data, nil
}

func (c Client) EmailsAndContacts(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/emails-and-contacts", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) PhonesEnricher(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/phones-enricher", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) AmazonProducts(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/amazon/products-v2", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) AmazonReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/amazon/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) YelpSearch(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/yelp-search", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) YelpReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/yelp/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) TripadvisorReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/tripadvisor-reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) AppStoreReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/appstore/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) YoutubeComments(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/youtube-comments", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) G2Reviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/g2/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) TrustpilotReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/trustpilot/reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GlassdoorReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/glassdoor/review", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) CapterraReviews(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/capterra-reviews", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) Geocoding(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/geocoding", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) ReverseGeocoding(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/reverse-geocoding", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) PhoneIdentityFinder(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/whitepages-phones", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) AddressScraper(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/whitepages-addresses", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) CompanyInsights(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/company-insights", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) ValidateEmails(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/email-validator", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) Trustpilot(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/trustpilot", parameters)
    return response["data"].([]interface{}), err
}


func (c Client) TrustpilotSearch(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/trustpilot", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) Similarweb(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/similarweb", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) CompanyWebsiteFinder(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/company-website-finder", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) YellowPages(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/yellowpages-search", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) postAPIRequest(path string, parameters map[string]interface{}) (Result, error) {
    httpClient := &http.Client{}
    fullUrl := ApiUrl + path

    payloadBytes, err := json.Marshal(parameters)
    if err != nil {
        return nil, errors.New("cannot encode payload")
    }

    req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(payloadBytes))
    req.Header.Add("X-API-KEY", c.ApiKey)
    req.Header.Add("Client", "Go SDK")
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")

    response, doError := httpClient.Do(req)
    if doError != nil {
        return nil, errors.New("cannot create request")
    }
    defer response.Body.Close()

    decoder := json.NewDecoder(response.Body)
    var result Result

    decodeError := decoder.Decode(&result)
    if decodeError != nil {
        return nil, errors.New("cannot decode")
    }

    return result, nil
}

func (c Client) PostAPIRequest(path string, parameters map[string]interface{}) (Result, error) {
    return c.postAPIRequest(path, parameters)
}

func (c Client) BusinessesSearch(
    filters map[string]interface{},
    limit int,
    includeTotal bool,
    cursor string,
    fields []string,
    asyncRequest bool,
    ui bool,
    webhook string,
    query string,
) (Result, error) {

    payload := map[string]interface{}{
        "filters":        map[string]interface{}{},
        "limit":          limit,
        "include_total":  includeTotal,
        "async":          asyncRequest,
        "ui":             ui,
    }

    if filters != nil {
        payload["filters"] = filters
    }
    if cursor != "" {
        payload["cursor"] = cursor
    }
    if fields != nil && len(fields) > 0 {
        payload["fields"] = fields
    }
    if webhook != "" {
        payload["webhook"] = webhook
    }

    if strings.TrimSpace(query) != "" {
        payload["query"] = query
    }

    res, err := c.postAPIRequest("/businesses", payload)
    if err != nil {
        return nil, err
    }

    if data, ok := res["data"]; ok {
        if dataObj, ok2 := data.(map[string]interface{}); ok2 {
            return dataObj, nil
        }
    }

    return res, nil
}

func (c Client) BusinessesIterSearch(
    filters map[string]interface{},
    limit int,
    fields []string,
    includeTotal bool,
) ([]interface{}, error) {

    cursor := ""
    all := make([]interface{}, 0)

    for {
        page, err := c.BusinessesSearch(
            filters,
            limit,
            includeTotal,
            cursor,
            fields,
            false,
            false,
            "",
            "",
        )
        if err != nil {
            return nil, err
        }

        itemsRaw, ok := page["items"]
        if !ok || itemsRaw == nil {
            break
        }

        items, ok := itemsRaw.([]interface{})
        if !ok || len(items) == 0 {
            break
        }

        all = append(all, items...)

        hasMore := false
        if hm, ok := page["has_more"].(bool); ok {
            hasMore = hm
        }

        nextCursor := ""
        if nc, ok := page["next_cursor"]; ok && nc != nil {
            if s, ok2 := nc.(string); ok2 {
                nextCursor = strings.TrimSpace(s)
            }
        }

        if !hasMore || nextCursor == "" {
            break
        }

        cursor = nextCursor
    }

    return all, nil
}

func (c Client) BusinessesGet(
    businessId string,
    fields []string,
    asyncRequest bool,
    ui bool,
    webhook string,
) (Result, error) {

    if strings.TrimSpace(businessId) == "" {
        return nil, errors.New("businessId is required")
    }

    params := map[string]string{}

    if fields != nil && len(fields) > 0 {
        params["fields"] = strings.Join(fields, ",")
    }
    if asyncRequest {
        params["async"] = "true"
    } else {
        params["async"] = "false"
    }
    if ui {
        params["ui"] = "true"
    }
    if webhook != "" {
        params["webhook"] = webhook
    }

    encodedId := url.PathEscape(businessId)

    res, err := c.getAPIRequest("/businesses/"+encodedId, params)
    if err != nil {
        return nil, err
    }

    if data, ok := res["data"]; ok {
        if dataObj, ok2 := data.(map[string]interface{}); ok2 {
            return dataObj, nil
        }

        if arr, ok2 := data.([]interface{}); ok2 && len(arr) > 0 {
            if firstObj, ok3 := arr[0].(map[string]interface{}); ok3 {
                return firstObj, nil
            }
        }
    }

    return res, nil
}
