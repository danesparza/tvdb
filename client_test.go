package tvdb_test

import (
	"testing"

	"github.com/danesparza/tvdb"
)

func TestTVDB_Login_ReturnsToken(t *testing.T) {
	//	Arrange
	request := tvdb.AuthRequest{}

	//	Act
	client := tvdb.TVDBClient{}
	response, err := client.Login(request)

	//	Assert
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}

	if response.Token == "" {
		t.Errorf("The token is blank, and shouldn't be")
	} else {
		t.Logf("Got a token back: %v", response.Token)
	}
}

func TestTVDB_SeriesSearch_ReturnsInformation(t *testing.T) {
	//	Arrange
	request := tvdb.SearchRequest{
		Name: "Looney"}

	//	Act
	client := tvdb.TVDBClient{Token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzIxMzUxMTUsImlkIjoidHZfcmVuYW1lciIsIm9yaWdfaWF0IjoxNDcyMDQ4NzE1fQ.Wu9C9CpfOmARfKpGU_00R-WDapReSIIgAvf2vLqNrTCF_uzgojaLMYgSJWbJAs0CjivvC6O11q0vslnZL7XVx6_36MILRMcxANGIm9Q6wCVkHDJOx4nsx2pCRHrwd2qQxJyFS-bye04dzcrxqEuzun-GFZpLThmziuSZArbd_s_loCbQfzB7ZJNcU0SbYFk-it28Z2byGvkmgrEpJ-RqEwTvcfSKsrhSf48akf64yjZ5xeScke6aoCZPWs39DmRD-_uvMFUtK2uX3nthbVWe4KTzMRF2SWUpYzH7910faH3wrB9PpjFl5FmIEJwRku0_c6XYKMQfg4ko3BhKrsDfag"}
	response, err := client.SeriesSearch(request)

	//	Assert
	if err != nil {
		t.Errorf("Error getting search results: %v", err)
	}

	if len(response.Data) == 0 {
		t.Errorf("There are no responses")
	}

	if response.Data[0].Id != 72514 {
		t.Errorf("Didn't get the series ID back that we expected")
	}
}
