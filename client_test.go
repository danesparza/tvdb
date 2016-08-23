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
