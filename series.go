package tvdb

// EpisodeRequest represents a request to get episode information
type EpisodeRequest struct {
	SeriesID int

	/* Query parameters */
	AiredEpisode int `json:"airedEpisode"`
	AiredSeason  int `json:"airedSeason"`
	DVDEpisode   int `json:"dvdEpisode"`
	DVDSeason    int `json:"dvdSeason"`
	IMDBId       int `json:"imdbId"`
}

// EpisodeResponse represents a reponse for episode information from the TVDB service
type EpisodeResponse struct {
	AbsoluteNumber     int             `json:"absoluteNumber"`
	AiredEpisodeNumber int             `json:"airedEpisodeNumber"`
	AiredSeason        int             `json:"airedSeason"`
	AiredSeasonID      int             `json:"airedSeasonID"`
	DVDEpisodeNumber   int             `json:"dvdEpisodeNumber"`
	DVDSeason          int             `json:"dvdSeason"`
	EpisodeName        string          `json:"episodeName"`
	FirstAired         string          `json:"firstAired"`
	ID                 int             `json:"id`
	Language           EpisodeLanguage `json:"language"`
	Overview           string          `json:"overview"`
}

// EpisodeLanguage represents the spoken language of the given episode
type EpisodeLanguage struct {
	EpisodeName     string `json:"episodeName"`
	EpisodeOverview string `json:"overview"`
}

// EpisodeResponseLinks represents the paging information for multiple episode responses
type EpisodeResponseLinks struct {
	FirstPage    int `json:"first"`
	LastPage     int `json:"last"`
	NextPage     int `json:"next"`
	PreviousPage int `json:"prev"`
}

// EpisodeResponses represents the list of responses to get episode information
type EpisodeResponses struct {
	Links EpisodeResponseLinks `json:"links"`
	Data  []EpisodeResponse    `json:"data"`
}
