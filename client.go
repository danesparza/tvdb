package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//	For more information on this service, please see the documentation site
//	located at: https://api.thetvdb.com/swagger

var (
	baseServiceURL = "https://api.thetvdb.com"
	apiKey         = "CA1E4A63116B1D87"
)

// Client is a service client to the TVDB service
type Client struct {
	ServiceURL string
	Token      string
	Language   string
}

// Login and get a bearer token
func (client *Client) Login(request AuthRequest) (AuthResponse, error) {
	retval := AuthResponse{}

	//	If the API key isn't set, just use the default:
	if request.APIKey == "" {
		request.APIKey = apiKey
	}

	//	If the API url isn't set, use the default:
	if client.ServiceURL == "" {
		client.ServiceURL = baseServiceURL
	}

	//	Set the API url
	apiURL := client.ServiceURL + "/login"

	//	Serialize our request to JSON:
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(&request)
	if err != nil {
		return retval, err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(requestBytes.String())

	//	Post the JSON to the api url
	res, err := http.Post(apiURL, "application/json", requestJSON)
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

// SeriesSearch search for a given TV series
func (client *Client) SeriesSearch(request SearchRequest) ([]SeriesInfo, error) {
	//	Create our return value
	retval := []SeriesInfo{}

	//	Initialize our client
	if err := client.initialize(); err != nil {
		return retval, err
	}

	//	Set the API url
	apiURL := client.ServiceURL + "/search/series"

	//	Construct our query
	u, err := url.Parse(apiURL)
	if err != nil {
		return retval, err
	}

	//	Update querystring parameters if necessary
	q := u.Query()

	if request.Name != "" {
		q.Set("name", request.Name)
	}

	if request.IMDBID != "" {
		q.Set("imdbId", request.IMDBID)
	}

	if request.Zap2ItID != "" {
		q.Set("zap2itId", request.Zap2ItID)
	}

	//	Encode the querystring
	u.RawQuery = q.Encode()

	//	Prep the response object
	searchResponse := SearchResponses{}

	//	Make the API call
	if err := client.makeAPIcall(u, &searchResponse); err != nil {
		return retval, err
	}
	retval = searchResponse.Data

	//	Return our response
	return retval, nil
}

// GetUpdated gets updated id's since a given unixtimestamp
func (client *Client) GetUpdated(request UpdatedRequest) ([]UpdatedResponse, error) {
	//	Create our return value
	retval := []UpdatedResponse{}

	//	Initialize our client
	if err := client.initialize(); err != nil {
		return retval, err
	}

	//	Set the API url
	apiURL := client.ServiceURL + "/updated/query"

	//	Construct our query
	u, err := url.Parse(apiURL)
	if err != nil {
		return retval, err
	}

	q := u.Query()

	if request.FromTime != 0 {
		q.Set("fromTime", strconv.FormatInt(request.FromTime, 10))
	}

	if request.ToTime != 0 {
		q.Set("toTime", strconv.FormatInt(request.ToTime, 10))
	}

	u.RawQuery = q.Encode()

	//	Prep the response object
	updatedResponses := UpdatedResponses{}

	//	Make the API call
	if err := client.makeAPIcall(u, &updatedResponses); err != nil {
		return retval, err
	}
	retval = updatedResponses.Data

	//	Return our response
	return retval, nil
}

// EpisodesForSeries searches for episodes in a given TV series
func (client *Client) EpisodesForSeries(request EpisodeRequest) ([]EpisodeResponse, error) {

	//	Create our return value
	retval := []EpisodeResponse{}

	//	Initialize our client
	if err := client.initialize(); err != nil {
		return retval, err
	}

	//	Set the API url
	apiURL := client.ServiceURL + fmt.Sprintf("/series/%v/episodes", request.SeriesID)

	//	TODO: If we have query options set on our request, update the url to use:

	//	Construct our query
	u, err := url.Parse(apiURL)
	if err != nil {
		return retval, err
	}

	//	Start with page 1
	for currentPage, lastPage := 1, 1; currentPage <= lastPage; currentPage++ {
		q := u.Query()
		q.Set("page", strconv.Itoa(currentPage))
		u.RawQuery = q.Encode()

		//	Prep the response object
		episodeResponses := EpisodeResponses{}

		//	Make the API call
		if err := client.makeAPIcall(u, &episodeResponses); err != nil {
			return retval, err
		}

		//	Append our results to the return value
		retval = append(retval, episodeResponses.Data...)

		//	Update the last page variable:
		lastPage = episodeResponses.Links.LastPage
	}

	//	Return our response
	return retval, nil
}

// makeAPIcall uses the url and model information to make the HTTP api call and deserialize the result
func (client *Client) makeAPIcall(u *url.URL, model interface{}) error {

	//	Create the request:
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	//	Set our headers:
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.Token)
	if client.Language != "" {
		req.Header.Set("Accept-Language", client.Language)
	}

	//	Make the request:
	res, err := httpClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Call not successful: %v", res.Status)
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(model)
	if err != nil {
		return err
	}

	return nil
}

// initialize authenticates & gets a bearer token and ensures the service url is set correctly
func (client *Client) initialize() error {
	//	If we don't have a token, get one first:
	if client.Token == "" {

		_, err := client.Login(AuthRequest{})
		if err != nil {
			return fmt.Errorf("Problem authenticating during search: %v", err)
		}
	}

	//	If the API url isn't set, use the default:
	if client.ServiceURL == "" {
		client.ServiceURL = baseServiceURL
	}

	return nil
}

// GetSeriesActors gets all available actors for a given show id
func (client *Client) GetSeriesActors(request SeriesRequest) ([]SeriesActorResponse, error) {
	//	Create our return value
	retval := []SeriesActorResponse{}

	//	Initialize our client
	if err := client.initialize(); err != nil {
		return retval, err
	}

	//	Set the API url
	apiURL := client.ServiceURL + fmt.Sprintf("/series/%v/actors", request.SeriesID)

	//	Construct our query
	u, err := url.Parse(apiURL)
	if err != nil {
		return retval, err
	}

	q := u.Query()

	u.RawQuery = q.Encode()

	//	Prep the response object
	object := SeriesActorResponses{}

	//	Make the API call
	if err := client.makeAPIcall(u, &object); err != nil {
		return retval, err
	}
	retval = object.Data

	//	Return our response
	return retval, nil
}

// GetSeriesImages gets images for a given show id and image type, if no KeyType is given it defaults to poster
func (client *Client) GetSeriesImages(request SeriesImageQueryRequest) ([]SeriesImageQueryResponse, error) {
	//	Create our return value
	retval := []SeriesImageQueryResponse{}

	//	Initialize our client
	if err := client.initialize(); err != nil {
		return retval, err
	}

	//	Set the API url
	apiURL := client.ServiceURL + fmt.Sprintf("/series/%v/images/query", request.SeriesID)

	//	Construct our query
	u, err := url.Parse(apiURL)
	if err != nil {
		return retval, err
	}

	q := u.Query()

	if request.KeyType != "" {
		q.Set("keyType", request.KeyType)
	} else {
		// make sure we have a keyType set!
		q.Set("keyType", "poster")
	}

	if request.Resulution != "" {
		q.Set("resolution", request.Resulution)
	}

	if request.SubKey != "" {
		q.Set("subKey", request.SubKey)
	}

	u.RawQuery = q.Encode()

	//	Prep the response object
	object := SeriesImageQueryResponses{}

	//	Make the API call
	if err := client.makeAPIcall(u, &object); err != nil {
		return retval, err
	}
	retval = object.Data

	//	Return our response
	return retval, nil
}
