package outscraper

import (
    "fmt"
    "net/url"
    "net/http"
    "encoding/json"
    "errors"
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
    response, err := c.getAPIRequest("/google-search-v2", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleSearchV3(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/google-search-v3", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsSearchV2(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/search-v2", parameters)
    return response["data"].([]interface{}), err
}

func (c Client) GoogleMapsReviewsV3(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/reviews-v3", parameters)
    return response["data"].([]interface{}), err
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
