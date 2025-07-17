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

func (c Client) GoogleMapsDirections(parameters map[string]string) ([]interface{}, error) {
    parameters["async"] = "false"
    response, err := c.getAPIRequest("/maps/directions", parameters)
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