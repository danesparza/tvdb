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
		Name: "Looney Tunes"}

	//	Act
	client := tvdb.TVDBClient{}
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
