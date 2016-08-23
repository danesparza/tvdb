package tvdb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

var (
	baseServiceUrl string = "https://api.thetvdb.com"
	apiKey         string = "CA1E4A63116B1D87"
)

type TVDBClient struct {
	ServiceUrl string
}

func (client TVDBClient) Login(request AuthRequest) (AuthResponse, error) {
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

	//	Return our response
	return retval, nil
}
