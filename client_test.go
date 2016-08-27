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
	matches, err := client.SeriesSearch(request)

	//	Assert
	if err != nil {
		t.Errorf("Error getting search results: %v", err)
	}

	if len(matches) == 0 {
		t.Errorf("There are no matches")
	}

	if matches[0].Id != 72514 {
		t.Errorf("Didn't get the series ID back that we expected")
	}
}

func TestTVDB_EpisodesForSeries_ReturnsInformation(t *testing.T) {
	//	Arrange
	request := tvdb.EpisodeRequest{
		SeriesId: 72514}

	//	Act
	client := tvdb.TVDBClient{}
	response, err := client.EpisodesForSeries(request)

	//	Assert
	if err != nil {
		t.Errorf("Error getting search results: %v", err)
	}

	if len(response) == 0 {
		t.Errorf("There are no responses")
	} else {
		t.Logf("Got %v episodes back", len(response))
	}

	if response[0].Id != 5657563 {
		t.Errorf("Didn't get the episode ID back that we expected")
	}
}

func TestTVDB_EpisodesForSeries_ReturnsExpectedCount(t *testing.T) {
	//	Arrange
	request := tvdb.EpisodeRequest{
		SeriesId: 305288}

	//	Act
	client := tvdb.TVDBClient{}
	response, err := client.EpisodesForSeries(request)

	//	Assert
	if err != nil {
		t.Errorf("Error getting search results: %v", err)
	}

	if len(response) != 8 {
		t.Errorf("8 episodes expected, but got %v instead", len(response))
	} else {
		t.Logf("Got %v episodes back", len(response))
	}

	if response[0].Id != 5468124 {
		t.Errorf("Didn't get the episode ID back that we expected")
	}
}

func TestTVDB_EpisodesForSeries_CanMap(t *testing.T) {
	//	Arrange
	request := tvdb.EpisodeRequest{
		SeriesId: 72514}

	//	Act
	client := tvdb.TVDBClient{}
	response, err := client.EpisodesForSeries(request)

	//	Assert
	if err != nil {
		t.Errorf("Error getting search results: %v", err)
	}

	if len(response) == 0 {
		t.Errorf("Didn't get any episodes")
	} else {
		t.Logf("Got %v episodes back", len(response))
	}

	//	Load up the map
	episodes := make(map[string]tvdb.EpisodeResponse)
	for _, episode := range response {
		episodes[episode.EpisodeName] = episode
	}

	t.Logf("Created a map with %v items in it", len(episodes))

	//	Check to see if the episode name exists
	//	and then get its season/episode number:
	episodeToFind := "Upswept Hare"
	if episode, ok := episodes[episodeToFind]; ok {
		if episode.AiredSeason != 1953 || episode.AiredEpisodeNumber != 7 {
			t.Errorf("The episode and season don't match what we expect. Expected s1953e7 - Found: s%ve%v", episode.AiredSeason, episode.AiredEpisodeNumber)
		}
	} else {
		t.Errorf("Didn't find the episode '%v'", episodeToFind)
	}
}
