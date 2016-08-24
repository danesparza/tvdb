package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	baseServiceUrl string = "https://api.thetvdb.com"
	apiKey         string = "CA1E4A63116B1D87"
)

type TVDBClient struct {
	ServiceUrl string
	Token      string
}

//	Login and get a bearer token
func (client *TVDBClient) Login(request AuthRequest) (AuthResponse, error) {
	retval := AuthResponse{}

	//	If the API key isn't set, just use the default:
	if request.ApiKey == "" {
		request.ApiKey = apiKey
	}

	//	If the API url isn't set, use the default:
	if client.ServiceUrl == "" {
		client.ServiceUrl = baseServiceUrl
	}

	//	Set the API url
	apiUrl := client.ServiceUrl + "/login"

	//	Serialize our request to JSON:
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(&request)
	if err != nil {
		return retval, err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(requestBytes.String())

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return retval, err
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return retval, err
	}

	//	Store the token:
	client.Token = retval.Token

	//	Return our response
	return retval, nil
}

//	Search for a given TV series
func (client *TVDBClient) SeriesSearch(request SearchRequest) (SearchResponses, error) {
	//	Create our return value
	retval := SearchResponses{}

	//	If we don't have a token, get one first:
	if client.Token == "" {

		_, err := client.Login(AuthRequest{})
		if err != nil {
			return retval, fmt.Errorf("Problem authenticating during search: %v", err)
		}
	}

	//	If the API url isn't set, use the default:
	if client.ServiceUrl == "" {
		client.ServiceUrl = baseServiceUrl
	}

	//	Set the API url
	apiUrl := client.ServiceUrl + "/search/series"

	//	Construct our query
	u, err := url.Parse(apiUrl)
	if err != nil {
		return retval, err
	}

	q := u.Query()

	if request.Name != "" {
		q.Set("name", request.Name)
	}

	if request.IMDBId != "" {
		q.Set("imdbId", request.IMDBId)
	}

	if request.Zap2ItId != "" {
		q.Set("zap2itId", request.Zap2ItId)
	}

	u.RawQuery = q.Encode()

	//	Create the request:
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return retval, err
	}

	//	Set our headers:
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.Token)

	//	Make the request:
	res, err := httpClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return retval, err
	}

	if res.StatusCode != 200 {
		return retval, fmt.Errorf("Call not successful: %v", res.Status)
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return retval, err
	}

	//	Return our response
	return retval, nil
}
