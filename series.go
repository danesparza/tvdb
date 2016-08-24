package tvdb

type EpisodeRequest struct {
	SeriesId int

	/* Query parameters */
	AiredEpisode int `json:"airedEpisode"`
	AiredSeason  int `json:"airedSeason"`
	DVDEpisode   int `json:"dvdEpisode"`
	DVDSeason    int `json:"dvdSeason"`
	IMDBId       int `json:"imdbId"`
}

type EpisodeResponse struct {
	AbsoluteNumber     int             `json:"absoluteNumber"`
	AiredEpisodeNumber int             `json:"airedEpisodeNumber"`
	AiredSeason        int             `json:"airedSeason"`
	AiredSeasonId      int             `json:"airedSeasonID"`
	DVDEpisodeNumber   int             `json:"dvdEpisodeNumber"`
	DVDSeason          int             `json:"dvdSeason"`
	EpisodeName        string          `json:"episodeName"`
	FirstAired         string          `json:"firstAired"`
	Id                 int             `json:"id`
	Language           EpisodeLanguage `json:"language"`
	Overview           string          `json:"overview"`
}

type EpisodeLanguage struct {
	EpisodeName     string `json:"episodeName"`
	EpisodeOverview string `json:"overview"`
}

type EpisodeResponseLinks struct {
	FirstPage    int `json:"first"`
	LastPage     int `json:"last"`
	NextPage     int `json:"next"`
	PreviousPage int `json:"prev"`
}

type EpisodeResponses struct {
	Links EpisodeResponseLinks `json:"links"`
	Data  []EpisodeResponse    `json:"data"`
}
