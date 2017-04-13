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
	ID                 int             `json:"id"`
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

// SeriesRequest is used to request additional series information
type SeriesRequest struct {
	SeriesID int `json:"id"`
}

// SeriesActorResponse contains information about a single actor
type SeriesActorResponse struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ImageAdded  string `json:"imageAdded"`
	ImageAuthor int    `json:"imageAuthor"`
	LastUpdated string `json:"lastUpdated"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	SeriesID    int    `json:"seriesId"`
	SortOrder   int    `json:"sortOrder"`
}

// SeriesActorResponses is the response of the api when asking for authors
type SeriesActorResponses struct {
	Data []SeriesActorResponse `json:"data"`
}

// SeriesImageQueryRequest used to query images for a given series and type
type SeriesImageQueryRequest struct {
	SeriesID   int    `json:"id"`
	KeyType    string `json:"keyType"`
	Resulution string `json:"resulution"`
	SubKey     string `json:"subKey"`
}

// SeriesImageQueryResponse one single image response
type SeriesImageQueryResponse struct {
	FileName    string             `json:"fileName"`
	ID          int                `json:"id"`
	KeyType     string             `json:"keyType"`
	LanguageID  int                `json:"languageId"`
	RatingsInfo map[string]float64 `json:"ratingsInfo"`
	Resolution  string             `json:"resulution"`
	SubKey      string             `json:"subKey"`
	Thumbnail   string             `json:"thumbnail"`
}

// SeriesImageQueryResponses is the response of the api when asking for images
type SeriesImageQueryResponses struct {
	Data []SeriesImageQueryResponse `json:"data"`
}
